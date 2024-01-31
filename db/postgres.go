package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/boomer-goten/nats-streaming-test/model"
	_ "github.com/lib/pq"
)

const (
	ConnStr      = "user=alex password=45hekmrf dbname=wb_base sslmode=disable"
	DatabaseName = "postgres"
)

type DataBaseStorage struct {
	db *sql.DB
}

func (dbs *DataBaseStorage) Open(DataBaseName string, ConnStr string) error {
	var err error = nil
	dbs.db, err = sql.Open(DataBaseName, ConnStr)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (dbs *DataBaseStorage) Close() {
	dbs.db.Close()
}

func (dbs *DataBaseStorage) InsertOrder(model *model.Order) error {
	tx, err := dbs.db.Begin()
	if err != nil {
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Order_info(order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardKey, sm_id, date_created, off_shard) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		model.OrderUID, model.TrackNumber, model.Entry, model.Locale, model.InternalSignature,
		model.CustomerID, model.DeliveryService, model.ShardKey, model.SMID, model.DateCreated, model.OOFShard)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Delivery(order_uid, name, phone, zip, city, address, region, email) VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		model.OrderUID, model.Delivery.Name, model.Delivery.Phone, model.Delivery.Zip, model.Delivery.City,
		model.Delivery.Address, model.Delivery.Region, model.Delivery.Email)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Payment(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		model.Payment.Transaction, model.Payment.RequestID, model.Payment.Currency, model.Payment.Provider, model.Payment.Amount,
		model.Payment.PaymentDt, model.Payment.Bank, model.Payment.DeliveryCost, model.Payment.GoodsTotal, model.Payment.CustomFee, model.OrderUID)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	for _, value := range model.Items {
		_, err = tx.Exec("INSERT INTO Items(chrt_id, transaction, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			value.ChrtID, model.Payment.Transaction, value.Price, value.RID, value.Name,
			value.Sale, value.Size, value.TotalPrice, value.NMID, value.Brand, value.Status)
		if err != nil {
			tx.Rollback()
			return errors.New("не удалось вставить данные в базу данных")
		}
	}
	err = tx.Commit()
	if err != nil {
		return errors.New("не удалось вставить данные в базу данных")
	}
	return err
}

func (dbs *DataBaseStorage) GetOrders() (map[string]model.Order, error) {
	rows, err := dbs.db.Query("SELECT * FROM Order_info JOIN Delivery on Order_info.order_uid = Delivery.order_uid JOIN Payment on Order_info.order_uid = Payment.order_uid")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := make(map[string]model.Order)
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature,
			&order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SMID, &order.DateCreated, &order.OOFShard, &order.OrderUID, &order.Delivery.Name,
			&order.Delivery.Phone, &order.Delivery.Zip, &order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
			&order.Delivery.Email, &order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency, &order.Payment.Provider,
			&order.Payment.Amount, &order.Payment.PaymentDt, &order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal,
			&order.Payment.CustomFee, &order.OrderUID); err != nil {
			return orders, err
		}
		query := "SELECT chrt_id, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items WHERE items.transaction = $1"
		item_rows, err := dbs.db.Query(query, order.Payment.Transaction)
		if err != nil {
			return nil, err
		}
		var items []model.Item
		for item_rows.Next() {
			var item model.Item
			if err := item_rows.Scan(&item.ChrtID, &item.Price, &item.RID, &item.Name,
				&item.Sale, &item.Size, &item.TotalPrice, &item.NMID, &item.Brand, &item.Status); err != nil {
				return orders, err
			}
			items = append(items, item)
		}
		order.Items = items
		orders[order.OrderUID] = order
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}
	return orders, nil
}

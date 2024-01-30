package db

import (
	"context"
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

func (dbs *DataBaseStorage) InsertOrder(model *model.Model) error {
	tx, err := dbs.db.Begin()
	if err != nil {
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Order_info(order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardKey, sm_id, date_created, off_shard) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		model.Order.OrderUid, model.Order.TrackNumber, model.Order.Entry, model.Order.Locale, model.Order.InternalSignature,
		model.Order.CustomerId, model.Order.DeliveryService, model.Order.ShardKey, model.Order.SmID, model.Order.DateCreated, model.Order.OffShard)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Delivery(order_uid, name, phone, zip, city, address, region, email) VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		model.Deliver.OrderUid, model.Deliver.Name, model.Deliver.Phone, model.Deliver.Zip, model.Deliver.City,
		model.Deliver.Address, model.Deliver.Region, model.Deliver.Email)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Payment(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		model.Pay.Transaction, model.Pay.RequestId, model.Pay.Currency, model.Pay.Provider, model.Pay.Amount,
		model.Pay.PaymentDt, model.Pay.Bank, model.Pay.DeliveryCost, model.Pay.GoodsTotal, model.Pay.CustomFee, model.Pay.OrderUid)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	_, err = tx.Exec("INSERT INTO Items(chrt_id, transaction, price, rid, name, sale, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		model.Item.ChrtId, model.Item.Transaction, model.Item.Price, model.Item.Rid, model.Item.Name,
		model.Item.Sale, model.Item.TotalPrice, model.Item.NmId, model.Item.Brand, model.Item.Status)
	if err != nil {
		tx.Rollback()
		return errors.New("не удалось вставить данные в базу данных")
	}
	err = tx.Commit()
	if err != nil {
		return errors.New("не удалось вставить данные в базу данных")
	}
	return err
}

func (dbs *DataBaseStorage) GetOrders(ctx context.Context) (map[string]model.Model, error) {
	rows, err := dbs.db.Query("SELECT * FROM Order_info JOIN Delivery on Order_info.order_uid = Delivery.order_uid JOIN Payment on Order_info.order_uid = Payment.order_uid JOIN items on items.transaction = Payment.transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders map[string]model.Model
	for rows.Next() {
		var order model.Model
		if err := rows.Scan(); err != nil {
			return orders, err
		}
		orders[order.Order.OrderUid] = order
	}
	if err = rows.Err(); err != nil {
		return orders, err
	}
	return orders, nil
}

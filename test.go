package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type OrderInfo struct {
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          int       `json:"shard_key"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OffShard          int       `json:"off_shard"`
}

type Delivery struct {
	OrderUid string `json:"order_uid"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Zip      string `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float32 `json:"amount"`
	PaymentDt    int     `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float32 `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    int     `json:"custom_fee"`
	OrderUid     string  `json:"order_uid"`
}

type Items struct {
	ChrtId      int     `json:"chrt_id"`
	Transaction string  `json:"transaction"`
	Price       float32 `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        float32 `json:"sale"`
	Size        string  `json:"size"`
	TotalPrice  float32 `json:"total_price"`
	NmId        int     `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}

type Model struct {
	Order    OrderInfo `json:"order"`
	Delivery Delivery  `json:"delivery"`
	Payment  Payment   `json:"payment"`
	Items    []Items   `json:"items"`
}

func main() {
	// Чтение данных из JSON-файла
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	// Создание экземпляра структуры Model
	var model Model

	// Преобразование JSON-данных в структуру Model
	err = json.Unmarshal(data, &model)
	if err != nil {
		fmt.Println("Ошибка при преобразовании JSON:", err)
		return
	}

	// Вывод данных
	fmt.Println("Order UID:", model.Order.OrderUid)
	fmt.Println("Track Number:", model.Order.TrackNumber)
	fmt.Println("Delivery OrderUID:", model.Delivery.OrderUid)
	fmt.Println("Delivery Name:", model.Delivery.Name)
	fmt.Println("Delivery Phone:", model.Delivery.Phone)
	fmt.Println("Delivery Zip:", model.Delivery.Zip)
	fmt.Println("Delivery City:", model.Delivery.City)
	fmt.Println("Delivery Address:", model.Delivery.Address)
	fmt.Println("Delivery Region:", model.Delivery.Region)
	fmt.Println("Delivery Email:", model.Delivery.Email)
	fmt.Println("Payment Transaction:", model.Payment.Transaction)
	fmt.Println("Items Count:", len(model.Items))
}

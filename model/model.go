package model

import (
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
	Zip      int    `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

type Payment struct {
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"requset_id"`
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
	TotalPrice  float32 `json:"total_price"`
	NmId        int     `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}

type Model struct {
	Order   OrderInfo
	Deliver Delivery
	Pay     Payment
	Item    Items
}

package db

import (
	"context"
	// "wb_l0/read_Data/model"
)

type DataBase interface {
	Open() error
	Close()
	// Restore() error
	InsertOrder(ctx context.Context, data model.Model) error
	GetOrders(ctx context.Context, orderId string) (map[string]model.Model, error)
}

var implementation DataBase

func Open() error {
	return implementation.Open()
}

func Close() {
	implementation.Close()
}

// func Restore() error {
// 	return implementation.Restore()
// }

func InsertOrder(ctx context.Context, data model.Model) error {
	return implementation.InsertOrder(ctx, data)
}

func GetOrders(ctx context.Context, orderId string) (map[string]model.Model, error) {
	return implementation.GetOrders(ctx, orderId)
}

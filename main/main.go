package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boomer-goten/nats-streaming-test/cache"
	"github.com/boomer-goten/nats-streaming-test/db"
	"github.com/boomer-goten/nats-streaming-test/model"
	"github.com/boomer-goten/nats-streaming-test/server"
	subscriber "github.com/boomer-goten/nats-streaming-test/sub"
	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

const (
	cluster_id = "test-cluster"
	channel    = "flow"
	client_id  = "sub"
)

func main() {
	var dbs db.DataBaseStorage
	err := dbs.Open(db.DatabaseName, db.ConnStr)
	if err != nil {
		log.Fatal("error database connection")
	}
	cacheMap := cache.NewCache()
	err = cacheMap.RestoreFromDB(&dbs)
	if err != nil {
		log.Fatal("error cache restore")
	}

	var subscriber subscriber.Sub
	subscriber.ConnectToStan(cluster_id, client_id)
	hand_msg := func(m *stan.Msg) {
		var receivedMsg model.Order
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := json.Unmarshal(m.Data, &receivedMsg)
		if err != nil {
			fmt.Println("Ошибка при преобразовании JSON:", err)
		}
		err = validate.Struct(receivedMsg)
		if err != nil {
			fmt.Println("Структура не прошла валидцаию:", err)
		} else {
			cacheMap.Add(receivedMsg.OrderUID, receivedMsg)
			dbs.InsertOrder(&receivedMsg)
		}
		fmt.Printf("Received: \n%s\n", string(m.Data))
	}
	subscriber.SubscribeToChannel(channel, hand_msg)
	server.RunServer(cacheMap)
	defer subscriber.CloseAll()
	defer cacheMap.Print()
	defer dbs.Close()
}

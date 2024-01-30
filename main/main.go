package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boomer-goten/nats-streaming-test/db"
	"github.com/boomer-goten/nats-streaming-test/model"
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
	sc, err := stan.Connect(cluster_id, client_id)
	if err != nil {
		log.Fatal("error_connect to stan")
	} else {
		fmt.Printf("Success Connection\n")
	}
	hand_msg := func(m *stan.Msg) {
		var receivedMsg model.Model
		err = json.Unmarshal(m.Data, &receivedMsg.Deliver)
		// if err == nil {
		dbs.InsertOrder(&receivedMsg)
		fmt.Printf("Received: %s\n", receivedMsg.Deliver.OrderUid)
		// }
		fmt.Printf("Received: %s\n", string(m.Data))
	}
	sub, err_two := sc.Subscribe(channel, hand_msg)
	defer sub.Unsubscribe()
	defer sub.Close()
	if err_two != nil {
		log.Fatal("error_subscribe")
	} else {
		fmt.Printf("Success Subscribe channel\n")
	}
	for {
		time.Sleep(time.Second)
	}
}

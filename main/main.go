package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

const (
	cluster_id = "test-cluster"
	channel    = "flow"
)

func main() {
	sc, err := stan.Connect(cluster_id, "sub")
	if err != nil {
		log.Fatal("error_connct")
	} else {
		fmt.Printf("Success Connection\n")
	}
	hand_msg := func(m *stan.Msg) {
		fmt.Printf("Received: %s\n", string(m.Data))
	}
	sub, err_two := sc.Subscribe("flow", hand_msg)
	defer sub.Close()
	defer sub.Unsubscribe()
	if err_two != nil {
		log.Fatal("error_subscribe")
	} else {
		fmt.Printf("Success Subscribe channel\n")
	}
	for {
		time.Sleep(time.Second)
	}
}

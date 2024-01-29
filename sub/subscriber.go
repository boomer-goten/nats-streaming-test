package subscriber

import (
	"fmt"
	"log"

	"github.com/nats-io/stan.go"
)

type Sub struct {
	cluster_id     string
	client_id      string
	channel        string
	StanConnection *stan.Conn
	StanSubscribe  *stan.Subscription
	handler        func(msg *stan.Msg)
}

func (sub *Sub) ConnectToStan(cluster_id, client_id string) (*stan.Conn, error) {
	sc, err := stan.Connect(cluster_id, client_id)
	if err != nil {
		log.Fatal("Connection_error")
	} else {
		sub.cluster_id = cluster_id
		sub.client_id = client_id
		sub.StanConnection = &sc
		fmt.Printf("Success Connection\n")
	}
	return sub.StanConnection, err
}

func (sub *Sub) SubscribeToChannel(channel string, handler func(msg *stan.Msg)) (*stan.Subscription, error) {
	subscribe, err := (*(sub.StanConnection)).Subscribe(channel, handler)
	if err != nil {
		log.Fatal("error_subscribe")
	} else {
		sub.channel = channel
		sub.handler = handler
		fmt.Printf("Success Subscribe channel\n")
	}
	return &subscribe, err
}

func (sub *Sub) CloseSub() error {
	if sub.StanSubscribe != nil {
		return (*(sub.StanSubscribe)).Unsubscribe()
	}
	return nil
}

func (sub *Sub) Close() error {
	if sub.StanSubscribe != nil {
		return (*(sub.StanSubscribe)).Close()
	}
	return nil
}

func (sub *Sub) CloseAll() {
	sub.Close()
	sub.CloseSub()
}

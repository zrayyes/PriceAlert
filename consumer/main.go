package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "kafka:9092"
)

var ctx = context.Background()

func CreateTopic() {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		panic(err.Error())
	}
}

func Consume() {
	l := log.New(os.Stdout, "kafka reader: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  l,
	})
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}

func main() {
	CreateTopic()
	Consume()
}

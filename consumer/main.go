package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"github.com/zrayyes/PriceAlert/consumer/models"
)

const (
	topic         = "message-log"
	brokerAddress = "kafka:9092"
)

var ctx = context.Background()
var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
		var alert models.Alert
		json.Unmarshal(msg.Value, &alert)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", alert.Email, alert.Coin, alert.Price)
	}
}

func main() {
	CreateTopic()
	Consume()
}

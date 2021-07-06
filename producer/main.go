package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"github.com/zrayyes/PriceAlert/producer/models"
)

const (
	topic         = "message-log"
	brokerAddress = "kafka:9092"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var ctx = context.Background()

func CreateTopic() {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		panic(err.Error())
	}
}

func Produce() {
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	for {
		alert, _ := json.Marshal(&models.Alert{ID: 1, Email: "Blo558@gmail.com", Coin: "BTC", Price: 35650.20})
		err := w.WriteMessages(ctx, kafka.Message{
			Value: []byte(alert),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second * 10)
	}
}

func main() {
	CreateTopic()
	Produce()
}

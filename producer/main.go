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
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  log.New(os.Stdout, "kafka writer: ", 0),
	})

	for {
		var alert models.Alert
		models.DB.Take(&alert, 1)
		alertJSON, _ := json.Marshal(&alert)

		err := kafkaWriter.WriteMessages(ctx, kafka.Message{
			Value: alertJSON,
		})
		if err != nil {
			fmt.Println("could not write message " + err.Error())
		}

		time.Sleep(time.Second * 10)
	}
}

func main() {
	models.ConnectDataBase()
	CreateTopic()
	Produce()
}

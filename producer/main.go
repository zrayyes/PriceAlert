package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"github.com/zrayyes/PriceAlert/producer/api"
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
		coins := []string{"BTC"}
		currencies := []string{"USD"}
		results := api.GetCoinPrices(coins, currencies)
		for _, coin := range coins {
			for _, currency := range currencies {
				price := results[coin][currency]
				fmt.Printf("%s -> %s = %f\n", coin, currency, price)

				var alerts []models.Alert
				models.DB.Where("coin = ? AND active = true AND price <= ?", coin, price).Find(&alerts)

				for _, alert := range alerts {
					alertJSON, _ := json.Marshal(&alert)
					err := kafkaWriter.WriteMessages(ctx, kafka.Message{
						Value: alertJSON,
					})
					if err != nil {
						fmt.Println("could not write message " + err.Error())
					}

					models.DB.Model(&alert).Update("active", false)
				}
			}
		}
		time.Sleep(time.Second * 60)
	}
}

func main() {
	models.ConnectDataBase()
	CreateTopic()
	Produce()
}

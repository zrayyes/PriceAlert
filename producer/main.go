package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"github.com/zrayyes/PriceAlert/producer/api"
	"github.com/zrayyes/PriceAlert/producer/helpers"
	"github.com/zrayyes/PriceAlert/producer/models"
)

var topic = helpers.GetEnv("KAFKA_TOPIC", "message-log")
var brokerAddress = fmt.Sprintf(helpers.GetEnv("KAFKA_HOST", "kafka"), ":", helpers.GetEnv("KAFKA_PORT", "9092"))

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var ctx = context.Background()

func CreateTopic() {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		panic(err.Error())
	}
}

func getAlerts(coin string, currency string, price float64) []models.Alert {
	var alerts []models.Alert
	models.DB.Where("active = true AND coin = ? AND active = true AND currency = ? AND price_min <= ? AND price_max >= ?", coin, currency, price, price).Find(&alerts)
	return alerts
}

func prepareAlertMessage(alert models.Alert, price float64) []byte {
	alertEvent := models.AlertEvent{Email: alert.Email, Coin: alert.Coin, Currency: alert.Currency, Price: price}
	alertEventJSON, _ := json.Marshal(&alertEvent)
	return alertEventJSON
}

func disableAlert(alert models.Alert) {
	models.DB.Model(&alert).Update("active", false)
}

func Produce() {
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  log.New(os.Stdout, "kafka writer: ", 0),
	})

	for {
		coins := strings.Split("BTC,ETC,BNB,XRP,DOT,USDT,DOGE,AXS,BUSD,ADA", ",")
		currencies := strings.Split("USD,EUR,GBP,JPY", ",")
		results := api.GetCoinPrices(coins, currencies)
		for _, coin := range coins {
			for _, currency := range currencies {
				price := results[coin][currency]
				if price != 0 {
					fmt.Printf("%s -> %s = %f\n", coin, currency, price)

					for _, alert := range getAlerts(coin, currency, price) {
						err := kafkaWriter.WriteMessages(ctx, kafka.Message{
							Value: prepareAlertMessage(alert, price),
						})
						if err != nil {
							fmt.Println("could not write message " + err.Error())
						}
						disableAlert(alert)
					}
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

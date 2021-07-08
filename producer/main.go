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

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var ctx = context.Background()

// Kafka values
var topic = helpers.GetEnv("KAFKA_TOPIC", "message-log")
var brokerAddress = fmt.Sprint(helpers.GetEnv("KAFKA_HOST", "kafka"), ":", helpers.GetEnv("KAFKA_PORT", "9092"))

// Connect to the Kafka broker to create a topic
func CreateTopic() error {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// Get all active alerts that are within the price range from the database
func getAlerts(coin string, currency string, price float64) []models.Alert {
	var alerts []models.Alert
	models.DB.Where("active = true AND coin = ? AND active = true AND currency = ? AND price_min <= ? AND price_max >= ?", coin, currency, price, price).Find(&alerts)
	return alerts
}

// JSONify an AlertEvent
func prepareAlertMessage(alert models.Alert, price float64) []byte {
	alertEvent := models.AlertEvent{Email: alert.Email, Coin: alert.Coin, Currency: alert.Currency, Price: price}
	alertEventJSON, _ := json.Marshal(&alertEvent)
	return alertEventJSON
}

// Disable an alert in the database
func disableAlert(alert models.Alert) {
	models.DB.Model(&alert).Update("active", false)
}

// Start the Kafka producer
func Produce() {
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  log.New(os.Stdout, "kafka writer: ", 0),
	})

	for {
		// A list of all the currently supported crypto coins (should be moved to DB)
		coins := strings.Split("BTC,ETC,BNB,XRP,DOT,USDT,DOGE,AXS,BUSD,ADA", ",")
		// A list of all the currently supported currencies (should be moved to DB)
		currencies := strings.Split("USD,EUR,GBP,JPY", ",")

		// Fetch the prices from cryptocompare
		results := api.GetCoinPrices(coins, currencies)
		for _, coin := range coins {
			for _, currency := range currencies {
				// Get the exchange rate for a crypto coin
				price := results[coin][currency]
				if price != 0 {
					// Find all the alerts that match  this price
					for _, alert := range getAlerts(coin, currency, price) {
						// Send an event to the Kafka topic
						err := kafkaWriter.WriteMessages(ctx, kafka.Message{
							Value: prepareAlertMessage(alert, price),
						})
						if err != nil {
							fmt.Println("could not write message " + err.Error())
						} else {
							// Disabled an alert after sending it to Kafka
							disableAlert(alert)
						}
					}
				}
			}
		}
		// Sleep for 15 seconds
		time.Sleep(time.Second * 15)
	}
}

func main() {
	models.ConnectDataBase()
	models.SetupDatabase()

	connected := false
	for !connected {
		if err := CreateTopic(); err != nil {
			fmt.Println("Failed to connected to Kafka, retrying in 15 seconds ...")
			time.Sleep(time.Second * 15)
		} else {
			fmt.Println("Connected to Kafka.")
			connected = true
		}
	}

	Produce()
}

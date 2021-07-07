package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
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
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
		Logger:  log.New(os.Stdout, "kafka reader: ", 0),
	})

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		var alert models.AlertEvent
		json.Unmarshal(msg.Value, &alert)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", alert.Email, alert.Coin, alert.Price)

		c, err := smtp.Dial("mailhog:1025")
		if err != nil {
			log.Fatal(err)
		}

		if err := c.Mail("noreply@pricealert.com"); err != nil {
			log.Fatal(err)
		}
		if err := c.Rcpt(alert.Email); err != nil {
			log.Fatal(err)
		}

		wc, err := c.Data()
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintf(wc, "From: noreply@pricealert.com\nTo: %s\nSubject: %s Price Alert\n\n%s has reached price %f %s.",
			alert.Email, alert.Coin, alert.Coin, alert.Price, alert.Currency)
		if err != nil {
			log.Fatal(err)
		}
		err = wc.Close()
		if err != nil {
			log.Fatal(err)
		}

		err = c.Quit()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	CreateTopic()
	Consume()
}

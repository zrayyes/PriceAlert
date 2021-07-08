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
	"github.com/zrayyes/PriceAlert/price-alert/helpers"
)

var topic = helpers.GetEnv("KAFKA_TOPIC", "message-log")
var brokerAddress = fmt.Sprintf(helpers.GetEnv("KAFKA_HOST", "kafka"), ":", helpers.GetEnv("KAFKA_PORT", "9092"))
var group = helpers.GetEnv("KAFKA_GROUP", "my-group")

var ctx = context.Background()
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func CreateTopic() {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		panic(err.Error())
	}
}

func getAlertFromMessage(msg kafka.Message) models.AlertEvent {
	var alert models.AlertEvent
	json.Unmarshal(msg.Value, &alert)
	return alert
}

func getEmailBody(alert models.AlertEvent) string {
	// Sender
	body := "From: noreply@pricealert.com\n"
	// Reciever
	body += fmt.Sprintf("To: %s\n", alert.Email)
	// Subject line
	body += fmt.Sprintf("%s Price Alert\n\n", alert.Coin)
	// Body
	body += fmt.Sprintf("%s has reached price %f %s.\n", alert.Coin, alert.Price, alert.Currency)

	return body
}

func sendEmail(alert models.AlertEvent) error {
	smtpAddress := fmt.Sprintf(helpers.GetEnv("SMTP_HOST", "mailhog"), ":", helpers.GetEnv("SMTP_PORT", "1025"))
	c, err := smtp.Dial(smtpAddress)
	if err != nil {
		return err
	}

	if err := c.Mail("noreply@pricealert.com"); err != nil {
		return err
	}
	if err := c.Rcpt(alert.Email); err != nil {
		return err
	}

	wc, err := c.Data()
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(wc, getEmailBody(alert))
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	err = c.Quit()
	if err != nil {
		return err
	}
	return nil
}

func Consume() {
	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: group,
		Logger:  log.New(os.Stdout, "kafka reader: ", 0),
	})

	for {
		msg, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		alert := getAlertFromMessage(msg)

		if err := sendEmail(alert); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func main() {
	CreateTopic()
	Consume()
}

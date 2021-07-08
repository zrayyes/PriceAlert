package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"
	"github.com/zrayyes/PriceAlert/consumer/helpers"
	"github.com/zrayyes/PriceAlert/consumer/models"
)

var ctx = context.Background()
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Kafka values
var topic = helpers.GetEnv("KAFKA_TOPIC", "message-log")
var brokerAddress = fmt.Sprint(helpers.GetEnv("KAFKA_HOST", "kafka"), ":", helpers.GetEnv("KAFKA_PORT", "9092"))
var group = helpers.GetEnv("KAFKA_GROUP", "my-group")

// Connect to the Kafka broker to create a topic
func CreateTopic() error {
	_, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, 0)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// Return the Kafka event message as an AlertEvent Struct
func getAlertFromMessage(msg kafka.Message) models.AlertEvent {
	var alert models.AlertEvent
	json.Unmarshal(msg.Value, &alert)
	return alert
}

// Create the email body in a RFC 822 message format
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

// Send out an email to the address declared in the alert
func sendEmail(alert models.AlertEvent) error {
	smtpAddress := fmt.Sprint(helpers.GetEnv("SMTP_HOST", "mailhog"), ":", helpers.GetEnv("SMTP_PORT", "1025"))
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

	if err = wc.Close(); err != nil {
		return err
	}

	if err = c.Quit(); err != nil {
		return err
	}
	return nil
}

// Start the Kafka consumer
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
			fmt.Println("could not read message: ", err.Error())
			continue
		}
		alert := getAlertFromMessage(msg)

		if err := sendEmail(alert); err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func main() {
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
	Consume()
}

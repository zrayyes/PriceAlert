package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

func Produce() {
	i := 0

	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}

func main() {
	CreateTopic()
	Produce()
}

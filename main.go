package main

import (
	"fmt"
	"kafka/args"
	"kafka/kafka"
	"time"
)

func main() {
	var brokers = args.Get("-brokers")
	var topic = args.Get("-topic")
	var group = args.Get("-group-id")
	var config = kafka.NewConsumerConfig()
	if group != "" {
		config = config.GroupID(group)
	}
	var consumer = kafka.NewConsumer(brokers, topic, config)

	fmt.Println("Started ...")

	consumer.OnMessage(func(message kafka.Message) {
		fmt.Println(string(message.Value))
	})

	for {
		time.Sleep(1 * time.Hour)
	}
}

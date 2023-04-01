package main

import (
	"email_sender_microservice/internal"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"strings"
)

// TODO: add config, mongodb

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "emails",
		"auto.offset.reset": "smallest",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("failed to connect to Kafka: %v", err)
	}
	defer consumer.Close()

	err = consumer.Subscribe("emails", nil)
	if err != nil {
		log.Fatalf("failed to subscribe to Kafka topic: %v", err)
	}

	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("msg from kafka queue %s\n", e.Value)
			email := strings.Split(string(e.Value), " ")[0]
			service.SendEmail(email)
		case kafka.Error:
			log.Printf("error %v\n", e)
		}
	}

}

package main

import (
	"context"
	"email_sender_microservice/internal"
	"email_sender_microservice/pkg/client/mongodb"
	"email_sender_microservice/pkg/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"strings"
)

func main() {
	cfg := config.GetConfig()

	config := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Server,
		"group.id":          cfg.Kafka.GroupId,
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
	log.Println("kafka works great")

	var ctx = context.Background()
	db, err := mongodb.NewMongoClient(ctx, cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.Database)
	if err != nil {
		log.Fatalf("failed to connect to database %v", err)
	}
	collection := db.Collection(cfg.Mongo.Collection)
	srv := service.NewService(collection)

	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Printf("msg from kafka queue %s\n", e.Value)
			email := strings.Split(string(e.Value), " ")[0]
			emailType := strings.Split(string(e.Value), " ")[1]
			srv.SendEmail(ctx, email, emailType)
		case kafka.Error:
			log.Printf("error %v\n", e)
		}
	}

}

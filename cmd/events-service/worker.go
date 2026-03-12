package main

import (
	"context"
	"log"

	"github.com/atul-engineer/document-service/internal/cache"
	"github.com/atul-engineer/document-service/internal/event"
)

func main() {
	// Initialize Kafka consumer
	kafkaReader := event.InitKafkaConsumer()
	//defer kafkaReader.Close()
	// Consume document events in a separate goroutine

	redisClient := cache.ConnectRedis()
	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Println("redis disconnected")
		}
	}()
	event.ConsumeDocumentEvents(context.Background(), kafkaReader, redisClient)
}
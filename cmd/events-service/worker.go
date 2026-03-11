package main

import (
	"context"

	"github.com/atul-engineer/document-service/internal/event"
)

func main() {
	// Initialize Kafka consumer
	kafkaReader := event.InitKafkaConsumer()
	//defer kafkaReader.Close()
	// Consume document events in a separate goroutine
	event.ConsumeDocumentEvents(context.Background(), kafkaReader)
}
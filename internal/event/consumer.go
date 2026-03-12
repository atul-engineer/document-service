package event

import (
	"context"
	"encoding/json"
	"log"

	"github.com/atul-engineer/document-service/internal/document"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
)

var DocumentCacheKey = "documents_cache"

func InitKafkaConsumer() *kafka.Reader {
	// Implementation for initializing Kafka consumer
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
	})
	// reader.SetOffset(kafka.LastOffset)
	// reader.Close()
	return reader
}

func ConsumeDocumentEvents(ctx context.Context, reader *kafka.Reader, redisClient *redis.Client) {
	// Implementation for consuming document events from Kafka
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			continue
		}
		log.Printf("Received order message: %s", string(msg.Value))

		var document document.DocumentEvent
		err = json.Unmarshal(msg.Value, &document)
		if err != nil {
			log.Printf("failed to unmarshal document message: %v", err)
			continue
		}
		if document.EventType == "created" {
			// Process the document event (e.g., update read model, trigger workflows, etc.)
			log.Printf("Processing document event for document ID: %s", document.DocumentID.Hex())
			// Invalidate cache for documents list
			if err := redisClient.Del(ctx, DocumentCacheKey).Err(); err != nil {
				log.Printf("Failed to invalidate documents cache: %v", err)
			} else {
				log.Println("Documents cache invalidated successfully")
			}

		} else {
			log.Printf("Unknown event type: %s", document.EventType)
		}
		log.Printf("Processed document event for document ID: %s", document.DocumentID.Hex())
	}
}

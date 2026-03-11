package event

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/atul-engineer/document-service/internal/document"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/v2/bson"
)


func InitKafkaProducer() *kafka.Writer {
	writer := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
		Topic: "orders",
		Balancer: &kafka.LeastBytes{},
	}
	return writer
}

func PublishDocumentEvent(ctx context.Context, writer *kafka.Writer, documentID bson.ObjectID, eventType string) {
	// Implementation for publishing document events to Kafka
	event := document.DocumentEvent{
		DocumentID: documentID,
		EventType:  eventType,
		Timestamp:  time.Now().Unix(),
	}
	// Serialize event and publish to Kafka
	message, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	err = writer.WriteMessages(ctx, kafka.Message{
		Key: []byte(event.DocumentID.Hex()),
		Value: message,
		Time: time.Now(),
	})

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	log.Println("Document message sent successfully")
	writer.Close()
}
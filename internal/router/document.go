package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/atul-engineer/document-service/internal/document"
	"github.com/atul-engineer/document-service/internal/event"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var DocumentCacheKey = "documents_cache"

func CreateDocument(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation for creating a document
		//documentService := document.NewDocumentService(client)
		var doc document.Document
		if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		documentService := document.NewDocumentService(client)
		res, err := documentService.Insert(r.Context(), &doc)
		if err != nil {
			http.Error(w, "Failed to create document", http.StatusInternalServerError)
			return
		}
		// Publish document event to Kafka
		eventProducer := event.InitKafkaProducer()
		event.PublishDocumentEvent(r.Context(), eventProducer, res.InsertedID.(bson.ObjectID), "created")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func ListDocuments(client *mongo.Client, redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation for listing documents
		documentService := document.NewDocumentService(client)
		cachedDocuments, err := redisClient.Get(r.Context(), DocumentCacheKey).Result()
		if err == nil {
			// Cache hit
			var documents []document.Document
			if err := json.Unmarshal([]byte(cachedDocuments), &documents); err == nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(documents)
				log.Println("Cache hit for documents total:", len(documents))
				return
			}
		}
		// Cache miss or unmarshal error, fetch from MongoDB
		documents, err := documentService.List(r.Context())
		if err != nil {
			http.Error(w, "Failed to list documents", http.StatusInternalServerError)
			return
		}
		// Cache the result in Redis		
		docBytes, err := json.Marshal(documents)
		if err == nil {
			redisClient.Set(r.Context(), DocumentCacheKey, docBytes, time.Minute*60) // Cache for 1 hr
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(documents)
	}
}
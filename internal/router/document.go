package router

import (
	"encoding/json"
	"net/http"

	"github.com/atul-engineer/document-service/internal/document"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}

func ListDocuments(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementation for listing documents
	}
}
package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartServer(client *mongo.Client) *http.Server {
	router := chi.NewRouter()
	router.Route("/documents", func(r chi.Router) {
		r.Post("/", CreateDocument(client))
		r.Get("/", ListDocuments(client))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return &server
}
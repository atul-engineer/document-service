package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/atul-engineer/document-service/internal/registry"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func StartServer(client *mongo.Client, redisClient *redis.Client) *http.Server {
	router := chi.NewRouter()
	router.Route("/documents", func(r chi.Router) {
		r.Post("/", CreateDocument(client))
		r.Get("/", ListDocuments(client, redisClient))
	})

	var port int
	fmt.Print("Enter the port number to run the server on: ")
	_, err := fmt.Scanln(&port)
	if err != nil {
		fmt.Println("Invalid input. Defaulting to port ", port)
		port = 8080
	}
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	registry, err := registry.NewRegistry([]string{"localhost:2379"})
	if err != nil {
		fmt.Println("Failed to connect to etcd registry:", err)
	}
	addr := "localhost:" + strconv.Itoa(port)
	_ = registry.Register("document-service", "instance-1", addr)
	return &server
}
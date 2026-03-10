package main

import (
	"context"
	"fmt"
	"log"

	"github.com/atul-engineer/document-service/internal/db"
	"github.com/atul-engineer/document-service/internal/router"
)

func main() {
	ctx := context.Background()
	mongoClient := db.ConnectMongoDB()
	defer func() {		
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Println("mongo disconnected")
		}
	}()

	server := router.StartServer(mongoClient)
	fmt.Println("Server is running on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
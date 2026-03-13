package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/atul-engineer/document-service/internal/discovery"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})

	d := discovery.NewDiscovery(client)

	for {
		services, err := d.GetServices("document-service")
		if err != nil {
			log.Fatal(err)
		}

		log.Println(services)

		for _, srv := range services {
			url := "http://" + srv + "/documents"
			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}
			body, _ := io.ReadAll(resp.Body)
			log.Println(string(body))
		}
		time.Sleep(3 * time.Second)
	}
}
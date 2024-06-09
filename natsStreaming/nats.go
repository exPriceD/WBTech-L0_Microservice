package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/nats-io/stan.go"
)

func main() {
	filePath := filepath.Join("natsStreaming", "order.json")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	clusterID := "test-cluster"
	clientID := "publisher-client"
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer func(sc stan.Conn) {
		_ = sc.Close()
	}(sc)

	channel := "orders"
	err = sc.Publish(channel, data)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	log.Printf("Message published to channel %s: %s", channel, string(data))
}

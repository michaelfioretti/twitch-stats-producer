package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func produceMessages(brokerAddress string) {
	topic := "test-topic"
	partition := 0

	fmt.Println("Producing message")
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka broker: %v", err)
	}
	fmt.Println("Connected to Kafka broker")
	fmt.Print(conn)

	defer conn.Close()

	fmt.Print("Writing message")
	if conn != nil {
		conn.WriteMessages(
			kafka.Message{Value: []byte("one!")},
		)
	}

	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
	)

	fmt.Println("Produced message")
}

func consumeMessages(brokerAddress string) {
	topic := "test-topic"
	partition := 0

	fmt.Println("Broker address:", brokerAddress)

	conn, _ := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)

	if conn == nil {
		log.Fatalf("Failed to connect to Kafka broker")
	}

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println("Received message:", string(b))
	}

	batch.Close()
}

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the Kafka broker address from the KAFKA_BROKER environment variable
	// brokerAddress := os.Getenv("KAFKA_BROKER")

	// go produceMessages(brokerAddress)

	// consumeMessages(brokerAddress)
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkahelper"
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
	fmt.Println("Starting Kafka producer")
	kafkahelper.ValidateBaseTopics()

	brokerAddress := kafkahelper.GetBrokerAddress()
	topic := "my-topic"
	fmt.Println("Broker address:", brokerAddress)

	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, 1)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	// go produceMessages(brokerAddress)

	// consumeMessages(brokerAddress)
}

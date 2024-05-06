package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func produceMessages(brokerAddress string) {
	topic := "test-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("Hello World!")},
	)

	fmt.Println("Produced message")
}

func consumeMessages(brokerAddress string) {
	topic := "test-topic"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", brokerAddress, topic, partition)

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
	brokerAddress := "kafka:9092"
	go produceMessages(brokerAddress)
	consumeMessages(brokerAddress)
}
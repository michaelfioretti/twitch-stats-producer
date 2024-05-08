package main

import (
	"fmt"
	"log"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkahelper"
	"github.com/segmentio/kafka-go"
)

func produceMessages() {
	fmt.Println("Starting Kafka producer")

	conn := kafkahelper.SetUpKafkaConnection()
	go kafkahelper.ValidateBaseTopics()

	defer conn.Close()

	fmt.Println("Connected to Kafka broker. Now writing messages")

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	fmt.Println("Produced messages")

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	fmt.Println("Produced message")
}

func consumeMessages() {
	conn := kafkahelper.SetUpKafkaConnection()

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
	go produceMessages()

	consumeMessages()
}

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkahelper"
	"github.com/segmentio/kafka-go"
)

func produceMessages() {
	conn := kafkahelper.SetUpKafkaConnection()
	go kafkahelper.ValidateBaseTopics()

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
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
}

func consumeMessages() {
	conn := kafkahelper.SetUpKafkaConnection()

	defer conn.Close()

	conn.SetReadDeadline(time.Time{})

	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		fmt.Println("Received message:", string(msg.Value))
		// Process the received message here
	}
}

func main() {
	go func() {
		for {
			produceMessages()
			time.Sleep(5 * time.Second)
		}
	}()

	consumeMessages()
}

package main

import (
	"log"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkaconsumer"
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

func main() {
	go func() {
		for {
			produceMessages()
			time.Sleep(5 * time.Second)
		}
	}()

	kafkaconsumer.ReadMessages()
}

package kafkaconsumer

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func CreateKafkaConsumer() *kafka.Reader {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	readerConfig := GetKafkaConsumerConfig()
	return kafka.NewReader(readerConfig)
}

func ReadMessages(reader *kafka.Reader) {
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// TODO: #5
func GetKafkaConsumerConfig() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092", "localhost:39092"},
		Topic:     "my-topic",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	}
}

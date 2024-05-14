package kafkaconsumer

import (
	"context"
	"fmt"
	"log"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkahelper"
	"github.com/segmentio/kafka-go"
)

func CreateKafkaConsumer(topic string) *kafka.Reader {
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	readerConfig := GetKafkaConsumerConfig(topic)
	return kafka.NewReader(readerConfig)
}

func ReadMessages(topic string) {
	reader := CreateKafkaConsumer(topic)
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

// TODO: #5 change Partition to GroupID
func GetKafkaConsumerConfig(topic string) kafka.ReaderConfig {
	brokerAddresses := kafkahelper.GetBrokerAddresses()
	return kafka.ReaderConfig{
		Brokers:   brokerAddresses,
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	}
}

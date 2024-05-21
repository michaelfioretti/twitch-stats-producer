package kafkaproducer

import (
	"testing"

	"github.com/segmentio/kafka-go"
)

func TestInitKafkaProducer(t *testing.T) {
	// Test that the producer is initialized correctly
	topic := "test-topic"
	err := InitKafkaProducer(topic)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if _, ok := producers[topic]; !ok {
		t.Errorf("Expected producer to be initialized, but it was not")
	}

	// Test that the producer is not initialized twice
	err = InitKafkaProducer(topic)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test that an error is returned if GetKafkaProducerConfig fails
	topic = "fail-topic"
	err = InitKafkaProducer(topic)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func TestWriteDataToKafka(t *testing.T) {
	// Test that writing data to Kafka works correctly
	topic := "test-topic"
	messages := []kafka.Message{
		{
			Value: []byte("message 1"),
		},
		{
			Value: []byte("message 2"),
		},
	}

	InitKafkaProducer(topic)

	WriteDataToKafka(topic, messages)

	// Test that writing data to a non-existent topic returns an error
	topic = "non-existent-topic"
	messages = []kafka.Message{
		{
			Value: []byte("message 1"),
		},
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic, but none occurred")
		}
	}()

	WriteDataToKafka(topic, messages)
}

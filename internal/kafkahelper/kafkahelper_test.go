package kafkahelper

import (
	"errors"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestInitKafkaProducer(t *testing.T) {
	t.Run("Initialize the producer", func(t *testing.T) {
		err := InitKafkaProducer("test")

		if err != nil {
			t.Error("Error initializing the producer", err)
		}
		assert.Equal(t, nil, err)
	})

	t.Run("Return an error if no topic is passed in", func(t *testing.T) {
		err := InitKafkaProducer("")
		assert.Equal(t, err, errors.New("topic is empty or nil"))
	})

	t.Run("Returns nil if the producer is already initialized", func(t *testing.T) {
		err := InitKafkaProducer("test")
		assert.Equal(t, nil, err)

		err = InitKafkaProducer("test")
		assert.Equal(t, nil, err)
	})
}

func TestWriteDataToKafka(t *testing.T) {
	t.Run("Write data to Kafka", func(t *testing.T) {
		topic := "test"
		messages := []kafka.Message{
			{Value: []byte("message 1")},
			{Value: []byte("message 2")},
		}

		err := WriteDataToKafka(topic, messages)
		assert.Equal(t, nil, err)
	})

	t.Run("Return an error if topic is empty", func(t *testing.T) {
		topic := ""
		messages := []kafka.Message{
			{Value: []byte("message 1")},
			{Value: []byte("message 2")},
		}

		err := WriteDataToKafka(topic, messages)
		assert.Equal(t, err, errors.New("topic is empty or nil"))
	})
}

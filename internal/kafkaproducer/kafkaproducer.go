// This file contains the implementation of being a producer
// for the Kafka cluster. The package will primarily read the
// live data and push it into the cluster.

package kafkaproducer

import (
	"context"
	"log"
	"sync"

	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/segmentio/kafka-go"
)

var (
	producers       = make(map[string]*kafka.Writer)
	producerConfigs = make(map[string]kafka.WriterConfig)
	mutex           sync.Mutex
)

func InitKafkaProducer(topic string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := producers[topic]; ok {
		return nil
	}

	producerConfig, err := getKafkaProducerConfig(topic)
	if err != nil {
		return err
	}

	producer := kafka.NewWriter(producerConfig)
	producers[topic] = producer
	producerConfigs[topic] = producerConfig

	return nil
}

func WriteDataToKafka(topic string, messages []kafka.Message) {
	producer := producers[topic]
	for _, message := range messages {
		err := producer.WriteMessages(context.Background(), message)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
	}
}

func getKafkaProducerConfig(topic string) (kafka.WriterConfig, error) {
	brokerAddresses := kafkahelper.GetBrokerAddresses()

	return kafka.WriterConfig{
		Brokers:  brokerAddresses,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}, nil
}

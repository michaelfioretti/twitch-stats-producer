// This file contains the implementation of being a producer
// for the Kafka cluster. The package will primarily read the
// live data and push it into the cluster.

package kafkaproducer

import (
	"fmt"
	"log"

	"github.com/michaelfioretti/twitch-stats-producer/internal/utils/kafkahelper"
	"github.com/segmentio/kafka-go"
)

func WriteDataToKafka(producer *kafka.Writer, data []byte) {
	// TODO: Write data here
	fmt.Println("Write the data here!.....")
}

func CreateKafkaProducer(topic string) *kafka.Writer {
	producerConfig := GetKafkaProducerConfig(topic)
	log.Default().Println("Creating Kafka producer...")
	return kafka.NewWriter(kafka.WriterConfig(producerConfig))
}

func GetKafkaProducerConfig(topic string) kafka.WriterConfig {
	brokerAddresses := kafkahelper.GetBrokerAddresses()

	return kafka.WriterConfig{
		Brokers:  brokerAddresses,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

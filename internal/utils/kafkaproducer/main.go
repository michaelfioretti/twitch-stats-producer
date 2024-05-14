// This file contains the implementation of being a producer
// for the Kafka cluster. The package will primarily read the
// live data and push it into the cluster.

package kafkaproducer

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func WriteDataToKafka(producer *kafka.Writer, data []byte) {
	// TODO: Write data here
	fmt.Println("Write the data here!.....")
}

func CreateKafkaProducer() *kafka.Writer {
	producerConfig := GetKafkaProducerConfig()
	log.Default().Println("Creating Kafka producer...")
	return kafka.NewWriter(kafka.WriterConfig(producerConfig))
}

func GetKafkaProducerConfig() kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers:  []string{"localhost:29092", "localhost:39092"},
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

package kafkahelper

import (
	"context"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

var (
	producers       = make(map[string]*kafka.Writer)
	producerConfigs = make(map[string]kafka.WriterConfig)
	mutex           sync.Mutex
)

func InitKafkaProducer(topic string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if topic == "" {
		return errors.New("topic is empty or nil")
	}

	if _, ok := producers[topic]; ok {
		log.Infof("Producer already exists for topic %s", topic)
		return nil
	}

	producerConfig, err := getKafkaProducerConfig(topic)
	if err != nil {
		log.Fatal("failed to get producer config:", err)
		return err
	}

	producer := kafka.NewWriter(producerConfig)
	producers[topic] = producer
	producerConfigs[topic] = producerConfig

	return nil
}

func WriteDataToKafka(topic string, messages []kafka.Message) error {
	InitKafkaProducer(topic)

	producer := producers[topic]
	err := producer.WriteMessages(context.Background(), messages...)

	if err != nil {
		log.Errorf("failed to write messages: %v", err)
		return err
	}

	return nil
}

func getKafkaProducerConfig(topic string) (kafka.WriterConfig, error) {
	brokerAddresses := getBrokerAddresses()

	return kafka.WriterConfig{
		Brokers:          brokerAddresses,
		Topic:            topic,
		Balancer:         &kafka.RoundRobin{},
		CompressionCodec: kafka.Lz4.Codec(),
	}, nil
}

func getBrokerAddresses() []string {
	brokerAddresses := os.Getenv("KAFKA_BROKERS")
	addresses := strings.Split(brokerAddresses, ",")
	return addresses
}

func getAvailableTopics() []string {
	topicsStr := constants.KAFKA_TOPICS
	availableTopics := strings.Split(topicsStr, ",")
	return availableTopics
}

func createKafkaConnection() *kafka.Conn {
	brokerAddresses := getBrokerAddresses()
	// Use first one for simplicity
	conn, err := kafka.Dial("tcp", brokerAddresses[0])
	if err != nil {
		log.Fatalf("Error connecting to Kafka cluster: %v", err)
	}

	return conn
}

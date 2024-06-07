package kafkahelper

import (
	"context"
	"errors"
	"os"
	"strconv"
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

// Pulls all topics listed in the .env file and creates them in the associated
// Kafka cluster if they do not exist.
// func ValidateBaseTopics() {
// 	partitionCount, replicationCount := getPartitionAndReplicationCount()

// 	conn := createKafkaConnection()

// 	defer conn.Close()

// 	controller, err := conn.Controller()
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	defer controllerConn.Close()

// 	// Loop and create topic configs
// 	topics := getAvailableTopics()
// 	topicConfigs := []kafka.TopicConfig{}
// 	for i := range topics {
// 		topicConfigs = append(topicConfigs, kafka.TopicConfig{
// 			Topic:             topics[i],
// 			NumPartitions:     partitionCount,
// 			ReplicationFactor: replicationCount,
// 		})
// 	}

// 	err = controllerConn.CreateTopics(topicConfigs...)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }

func getAvailableTopics() []string {
	topicsStr := constants.KAFKA_TOPICS
	availableTopics := strings.Split(topicsStr, ",")
	return availableTopics
}

// func createKafkaConnection() *kafka.Conn {
// 	brokerAddresses := getBrokerAddresses()
// 	// Use first one for simplicity
// 	conn, err := kafka.Dial("tcp", brokerAddresses[0])
// 	if err != nil {
// 		log.Fatalf("Error connecting to Kafka cluster: %v", err)
// 	}

// 	return conn
// }

func getPartitionAndReplicationCount() (int, int) {
	partitionCountStr := os.Getenv("PARTITION_COUNT")
	replicationCountStr := os.Getenv("REPLICATION_COUNT")

	partitionCount, err := strconv.Atoi(partitionCountStr)
	if err != nil {
		log.Fatalf("Error converting PARTITION_COUNT to integer: %v", err)
	}

	replicationCount, err := strconv.Atoi(replicationCountStr)
	if err != nil {
		{
			log.Fatalf("Error converting REPLICATION_COUNT to integer: %v", err)
		}
	}

	return partitionCount, replicationCount
}

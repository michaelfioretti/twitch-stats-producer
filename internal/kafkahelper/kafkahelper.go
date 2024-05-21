package kafkahelper

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/utils"
	"github.com/segmentio/kafka-go"
)

// Create interface so that we can use dependency injection
// to mock these functions during tests
type KafkaHelper interface {
	GetBrokerAddresses() []string
	ValidateBaseTopics()
}

func GetBrokerAddresses() []string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	brokerAddresses := os.Getenv("KAFKA_BROKERS")
	addresses := strings.Split(brokerAddresses, ",")
	return addresses
}

// Pulls all topics listed in the .env file and creates them in the associated
// Kafka cluster if they do not exist.
func ValidateBaseTopics() {

	partitionCount, replicationCount := getPartitionAndReplicationCount()

	conn := createKafkaConnection()

	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		log.Fatal(err.Error())
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer controllerConn.Close()

	// Loop and create topic configs
	topics := getAvailableTopics()
	topicConfigs := []kafka.TopicConfig{}
	for i := range topics {
		topicConfigs = append(topicConfigs, kafka.TopicConfig{
			Topic:             topics[i],
			NumPartitions:     partitionCount,
			ReplicationFactor: replicationCount,
		})
	}

	fmt.Print("Creating topics...\n")

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getAvailableTopics() []string {
	topicsStr := constants.KAFKA_TOPICS
	availableTopics := strings.Split(topicsStr, ",")
	return availableTopics
}

func createKafkaConnection() *kafka.Conn {
	brokerAddresses := GetBrokerAddresses()
	// Use first one for simplicity
	conn, err := kafka.Dial("tcp", brokerAddresses[0])
	if err != nil {
		log.Fatalf("Error connecting to Kafka cluster: %v", err)
	}

	return conn
}

func getPartitionAndReplicationCount() (int, int) {
	partitionCountStr := utils.GetEnvVar("PARTITION_COUNT")
	replicationCountStr := utils.GetEnvVar("REPLICATION_COUNT")

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

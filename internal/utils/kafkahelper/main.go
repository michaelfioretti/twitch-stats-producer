package kafkahelper

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func GetBrokerAddress() string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv("KAFKA_BROKER")
}

func SetUpKafkaConnection() *kafka.Conn {
	brokerAddress := GetBrokerAddress()

	// Note: right now the topic and partition are hardcoded, but this will change
	conn, err := kafka.DialLeader(context.Background(), "tcp", brokerAddress, "my-topic", 0)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka broker: %v", err)
	}

	return conn
}

func ValidateBaseTopics() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the Kafka broker address from the KAFKA_BROKER environment variable
	brokerAddress := os.Getenv("KAFKA_BROKER")
	topic := "my-topic"

	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}

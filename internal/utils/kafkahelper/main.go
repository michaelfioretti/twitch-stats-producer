package kafkahelper

import (
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

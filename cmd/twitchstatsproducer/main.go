package main

import (
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaconsumer"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	"github.com/segmentio/kafka-go"
)

func produceMessages() {
	go kafkahelper.ValidateBaseTopics()

	msg1 := kafka.Message{Value: []byte("one!")}
	msg2 := kafka.Message{Value: []byte("two!")}
	msg3 := kafka.Message{Value: []byte("three!")}

	kafkaproducer.WriteDataToKafka("my-topic", []kafka.Message{msg1, msg2, msg3})
}

func main() {
	go func() {
		for {
			produceMessages()
			time.Sleep(5 * time.Second)
		}
	}()

	kafkaconsumer.ReadMessages("my-topic")
}

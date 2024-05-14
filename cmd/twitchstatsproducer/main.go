package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
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
	// Helper code to keep server running (for now!)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "eventsub.wss.twitch.tv", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	defer conn.Close()

	wsMessages := make(chan struct{})

	go twitchhelper.ListenToMessages(conn, wsMessages)

	// Channel management based on messages received from Twitch
	for {
		select {
		case <-wsMessages:
			return
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error:", err)
			}
			select {
			case <-wsMessages:
			}
			return
		}
	}
}

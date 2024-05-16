package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

	clientId, _ := twitchhelper.LoadTwitchKeys()

	// First, get the OAuth token
	err, oauthToken := twitchhelper.SendOauthRequest()

	if err != nil {
		log.Fatal("Error getting OAuth token:", err)
	}

	liveChannelsCount, err := twitchhelper.GetLiveChannelsCount(clientId, oauthToken.AccessToken)
	if err != nil {
		fmt.Printf("Error getting live channels count: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("There are currently %d live channels on Twitch.\n", liveChannelsCount)

	// u := url.URL{Scheme: "wss", Host: constants.TWITCH_PUBSUB_URL}
	// headers := http.Header{}
	// headers.Add("Authorization", "Bearer "+oauthToken.AccessToken)

	// log.Printf("Securely connecting to %s", u.String())

	// conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	// if err != nil {
	// 	log.Fatal("Dial error:", err)
	// }

	// defer conn.Close()

	// // Subscribe to a specific topic
	// subscribeMessage := `{"type": "LISTEN", "data": {"topics": ["` + topic + `"], "auth_token": "` + accessToken + `"}}`
	// err = conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
	// if err != nil {
	// 	return fmt.Errorf("writing subscribe message failed: %v", err)
	// }

	// fmt.Println("Connected to Twitch PubSub API. Waiting for messages...")

	// // Read messages from the connection
	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		return fmt.Errorf("error reading message: %v", err)
	// 	}
	// 	fmt.Printf("Received message: %s\n", message)
	// 	// Process the received message here (e.g., parse JSON and handle accordingly)
	// }

	// wsMessages := make(chan struct{})

	// go twitchhelper.ListenToMessages(conn, wsMessages)

	// // Channel management based on messages received from Twitch
	// for {
	// 	select {
	// 	case <-wsMessages:
	// 		return
	// 	case <-interrupt:
	// 		log.Println("Interrupt received, closing connection...")
	// 		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// 		if err != nil {
	// 			log.Println("Write close error:", err)
	// 		}
	// 		select {
	// 		case <-wsMessages:
	// 		}
	// 		return
	// 	}
	// }
}

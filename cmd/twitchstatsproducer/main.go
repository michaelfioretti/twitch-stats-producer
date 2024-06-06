package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
)

func main() {
	messageChannel := make(chan *models.TwitchMessage, 1000)

	log.Println("Starting Kafka producer")
	go kafkahelper.ValidateBaseTopics()

	client := twitchchatparser.CreateTwitchClient()
	log.Println("Got the client.....")
	twitchchatparser.SubscribeToTwitchChat(client)
	log.Println("Subscribing.....")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		twitchMessage := twitchchatparser.ParseTwitchMessage(message)
		messageChannel <- twitchMessage
	})

	// We will update the top 100 streamers every 5 minutes
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Updating top 100 streamers at this time: ", time.Now().Format("2006-01-02 15:04:05"))
				twitchchatparser.UpdateStreamerList(client)
			}
		}
	}()

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")
}

package main

import (
	"log"
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/shared"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	go kafkahelper.ValidateBaseTopics()

	shared.TwitchClient = twitchchatparser.CreateTwitchClient()
	twitchchatparser.SubscribeToTwitchChat()

	shared.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		twitchMessage := twitchchatparser.ParseTwitchMessage(message)
		shared.MessageChannel <- twitchMessage
	})

	go twitchchatparser.ProcessTwitchMessages()
	go shared.TwitchClient.Connect()

	http.ListenAndServe(":8080", nil)
}

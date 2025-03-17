package main

import (
	"net/http"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/mongodbhelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/shared"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
)

func main() {
	mongodbhelper.ConnectToMongoDb()
	defer mongodbhelper.DisconnectFromMongoDn()

	shared.TwitchClient = twitchchatparser.CreateTwitchClient()
	twitchchatparser.SubscribeToTwitchChat()

	shared.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		twitchMessage := twitchchatparser.ParseTwitchMessage(message)
		shared.MessageChannel <- twitchMessage
	})

	go mongodbhelper.ProcessTwitchMessages()
	go shared.TwitchClient.Connect()

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
	"github.com/segmentio/kafka-go"
)

func main() {
	go kafkahelper.ValidateBaseTopics()

	// First, get the OAuth token
	err, oauthToken := twitchhelper.SendOauthRequest()

	if err != nil {
		log.Fatal("Error getting OAuth token:", err)
	}

	token := "oauth:" + oauthToken.AccessToken

	topLivestreams, err := twitchhelper.GetTop100Livestreams(oauthToken.AccessToken)
	if err != nil {
		log.Fatalf("Error getting livestreams: %v\n", err)
	}

	doneChan := make(chan struct{})

	// Connect to Twitch IRC using your Twitch username and auth token.
	// Make sure you registered your app on the Twitch Developer Dashboard to obtain the auth token.
	client := twitch.NewClient(constants.TWITCH_USERNAME, token)
	client.Capabilities = []string{twitch.TagsCapability, twitch.CommandsCapability, twitch.MembershipCapability} // Customize which capabilities are sent

	streamerNames := make([]string, 100)

	for _, stream := range topLivestreams {
		streamerNames = append(streamerNames, stream.UserName)
	}

	fmt.Println("Streamer names: ", streamerNames)

	go client.Join(streamerNames...)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// Parse the incoming message.
		parsedMessage := twitchchatparser.ParseTwitchMessage(message)

		str, err := json.Marshal(parsedMessage)
		if err != nil {
			log.Printf("Error marshaling parsed message: %v\n", err)
			return
		}
		msg := kafka.Message{Value: str}
		go kafkaproducer.WriteDataToKafka("streamer_chat", []kafka.Message{msg})
	})

	client.Connect()

	<-doneChan
}

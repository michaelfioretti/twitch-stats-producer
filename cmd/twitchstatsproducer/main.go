package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
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

	for _, stream := range topLivestreams.Data {
		streamerNames = append(streamerNames, stream.UserName)
	}

	fmt.Println("Streamer names: ", streamerNames)

	go client.Join(streamerNames...)

	// Listen to all messages in the channel, and update
	// Kafka cluster in batches

	batch := make([]kafka.Message, 0)
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		// Parse the incoming message.
		parsedMessage := twitchchatparser.ParseTwitchMessage(message)

		twitchMessage, err := proto.Marshal(parsedMessage)
		if err != nil {
			log.Fatalf("Error marshaling message: %v\n", err)
			return
		}

		msg := kafka.Message{Value: twitchMessage}
		batch = append(batch, msg)

		if len(batch) == 100 {
			log.Println("Writing 100 more messages at this time: ", time.Now().Format("2006-01-02 15:04:05"))
			print(batch)
			go kafkaproducer.WriteDataToKafka("streamer_chat", batch)
			batch = make([]kafka.Message, 0)
		}
	})

	client.Connect()

	<-doneChan
}

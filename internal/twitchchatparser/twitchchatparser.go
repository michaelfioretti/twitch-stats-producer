package twitchchatparser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

// Using this channel to keep the method non blocking so we can update the streamers periodically
var messageChannel = make(chan *models.TwitchMessage)

func CreateTwitchClient() *twitch.Client {
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
		parsedMessage := parseTwitchMessage(message)

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

	return client
}

func CreateTwitchClientWithChannel() *twitch.Client {

}

func UpdateStreamerList(client *twitch.Client) {
	fmt.Println("Hey got here.....")
}

func parseTwitchMessage(message twitch.PrivateMessage) *models.TwitchMessage {
	// Extract the username from the message sender.
	username := message.User.Name

	// Get the message content.
	messageText := message.Message

	// Check if the message is an action (/me).
	isAction := strings.HasPrefix(messageText, "\x01ACTION ")
	if isAction {
		// Remove the action prefix to get the actual message.
		messageText = strings.TrimPrefix(messageText, "\x01ACTION ")
		messageText = strings.TrimSuffix(messageText, "\x01")
	}

	// Extract badges from the message tags.
	var badges []string
	badgeStr := message.Tags["badges"]
	if badgeStr != "" {
		badges = strings.Split(badgeStr, ",")
	}

	// Extract the number of bits cheered (if any).
	bitsCheered, _ := strconv.ParseInt(message.Tags["bits"], 10, 32)

	// Get the channel name from the message tags.
	channel := message.Tags["room-id"]

	subscribed, _ := strconv.Atoi(message.Tags["subscriber"])
	mod, _ := strconv.Atoi(message.Tags["mod"])

	// Return the structured data.
	return &models.TwitchMessage{
		Username:   username,
		Channel:    message.Channel,
		Message:    messageText,
		Badges:     badges,
		Bits:       int32(bitsCheered),
		Mod:        int32(mod),
		Subscribed: int32(subscribed),
		Color:      message.Tags["color"],
		RoomID:     "#" + channel,
	}
}

package twitchchatparser

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
)

func CreateTwitchClient() *twitch.Client {
	// First, get the OAuth token
	oauthToken := twitchhelper.SendOauthRequest()

	token := "oauth:" + oauthToken.AccessToken

	client := twitch.NewClient(constants.TWITCH_USERNAME, token)
	client.Capabilities = []string{twitch.TagsCapability, twitch.CommandsCapability, twitch.MembershipCapability} // Customize which capabilities are sent

	client.Connect()

	return client
}

// Used to get the initial top 100 streamers. Subsequent updates will be done in the UpdateStreamerList method
func SubscribeToTwitchChat(client *twitch.Client) {
	oauthToken := twitchhelper.SendOauthRequest()

	topLivestreams, err := twitchhelper.GetTop100Livestreams(oauthToken.AccessToken)
	if err != nil {
		log.Fatalf("Error getting livestreams: %v\n", err)
	}

	streamerNames := make([]string, 100)

	for _, stream := range topLivestreams.Data {
		streamerNames = append(streamerNames, stream.UserName)
	}

	go client.Join(streamerNames...)
}

func UpdateStreamerList(client *twitch.Client) {
	fmt.Println("Hey got here.....")
}

func ParseTwitchMessage(message twitch.PrivateMessage) *models.TwitchMessage {
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

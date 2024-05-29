package twitchchatparser

import (
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
)

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

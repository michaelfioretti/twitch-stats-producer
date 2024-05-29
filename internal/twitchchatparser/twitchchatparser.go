package twitchchatparser

import (
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

// TwitchMessage represents a formatted Twitch chat message.
type TwitchMessage struct {
	Username   string
	Channel    string
	Message    string
	Badges     []string
	Bits       int
	Mod        int
	Subscribed int
	Color      string
	RoomID     string
}

func ParseTwitchMessage(message twitch.PrivateMessage) TwitchMessage {
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
	bitsCheered, _ := strconv.Atoi(message.Tags["bits"])

	// Get the channel name from the message tags.
	channel := message.Tags["room-id"]

	subscribed, _ := strconv.Atoi(message.Tags["subscriber"])
	mod, _ := strconv.Atoi(message.Tags["mod"])

	// Return the structured data.
	return TwitchMessage{
		Username:   username,
		Channel:    message.Channel,
		Message:    messageText,
		Badges:     badges,
		Bits:       bitsCheered,
		Mod:        mod,
		Subscribed: subscribed,
		Color:      message.Tags["color"],
		RoomID:     "#" + channel,
	}
}

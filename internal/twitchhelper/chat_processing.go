package twitchhelper

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TwitchChatMessage struct {
	Timestamp    time.Time `json:"timestamp"`     // Precise message time (UTC)
	Channel      string    `json:"channel"`       // Channel where the message was sent
	Username     string    `json:"username"`      // Username of the sender
	UserID       string    `json:"user_id"`       // Unique Twitch ID of the sender
	DisplayName  string    `json:"display_name"`  // Display name (may differ from username)
	MessageText  string    `json:"message_text"`  // Actual text content of the message
	IsModerator  bool      `json:"is_moderator"`  // Whether the sender is a moderator
	IsSubscriber bool      `json:"is_subscriber"` // Whether the sender is a subscriber
	Bits         int       `json:"bits"`          // Number of bits used in a cheer (0 if not applicable)
	Emotes       []Emote   `json:"emotes"`        // List of emotes used in the message
	MessageType  string    `json:"message_type"`  // Type of message (e.g., "chat", "whisper", "action")
}

type Emote struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"` // Number of times this emote was used in the message
}

// ParseTwitchChatMessage takes a raw Twitch chat message string and returns a TwitchChatMessage struct
func ConvertToKafkaMessage(rawMessage string) (*TwitchChatMessage, error) {
	// Example rawMessage: "@badges=moderator/1;color=#0000FF;display-name=CoolStreamer;emotes=25:0-4,12-16/1902:6-10;id=13521e19-1698-44d2-83e2-12a8c6848379;mod=1;room-id=123456789;subscriber=0;tmi-sent-ts=1593245600000;turbo=0;user-id=138425719;user-type=mod :coolstreamer!coolstreamer@coolstreamer.tmi.twitch.tv PRIVMSG #coolstreamer :Hey guys! Welcome to the stream!"

	// Split the message into tags and the rest
	parts := strings.SplitN(rawMessage, " :", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid message format")
	}
	tagsStr := parts[0]
	messageText := parts[1]

	// Parse tags
	tags := make(map[string]string)
	for _, tag := range strings.Split(tagsStr[1:], ";") { // Skip initial '@'
		tagParts := strings.SplitN(tag, "=", 2)
		if len(tagParts) == 2 {
			tags[tagParts[0]] = tagParts[1]
		}
	}

	// Extract information
	timestamp, _ := strconv.ParseInt(tags["tmi-sent-ts"], 10, 64)
	isModerator := tags["mod"] == "1"
	isSubscriber := tags["subscriber"] == "1"
	bits, _ := strconv.Atoi(tags["bits"])
	channel := strings.TrimPrefix(tags["room-id"], "#") // Extract channel without '#'
	displayName := tags["display-name"]
	username := strings.Split(messageText, "!")[0]
	userID := tags["user-id"]

	// Parse emotes (if present)
	var emotes []Emote
	if emoteTag, exists := tags["emotes"]; exists {
		emoteParts := strings.Split(emoteTag, "/")
		for _, emotePart := range emoteParts {
			emoteInfo := strings.Split(emotePart, ":")
			emoteID := emoteInfo[0]
			emoteName := strings.Split(emoteInfo[1], "-")[0]                   // Take the first occurrence of the emote
			emoteCount, _ := strconv.Atoi(strings.Split(emoteInfo[1], ",")[0]) // Take the first count
			emotes = append(emotes, Emote{ID: emoteID, Name: emoteName, Count: emoteCount})
		}
	}

	// Create TwitchChatMessage struct
	return &TwitchChatMessage{
		Timestamp:    time.Unix(0, timestamp*int64(time.Millisecond)), // Convert to time.Time
		Channel:      channel,
		Username:     username,
		UserID:       userID,
		DisplayName:  displayName,
		MessageText:  messageText,
		IsModerator:  isModerator,
		IsSubscriber: isSubscriber,
		Bits:         bits,
		Emotes:       emotes,
		MessageType:  "chat", // Assuming PRIVMSG is a chat message (adjust if needed)
	}, nil
}

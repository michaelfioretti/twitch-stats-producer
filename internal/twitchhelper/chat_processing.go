package twitchhelper

import (
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

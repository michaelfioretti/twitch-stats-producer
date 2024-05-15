// Main helper to connect to Twitch events
package twitchhelper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/michaelfioretti/twitch-stats-producer/internal/utils"
)

const (
	twitchWelcomeMessage = "session_welcome"
)

// type TwitchMessage struct {
// 	Metadata struct `json:"type"`
// 	Payload string `json:"data"`
// }

type TwitchMessageRequest struct {
	Type      string `json:"type"`
	Version   string `json:"version"`
	Condition struct {
		BroadcasterUserId string `json:"broadcaster_user_id"`
		ModeratorUserId   string `json:"moderator_user_id"`
	}
}

func main() {
	fmt.Println("Hello from Twitch Helper!")
}

func LoadTwitchKeys() (string, string) {
	// Load Twitch keys from .env file
	clientId := utils.GetEnvVar("TWITCH_CLIENT_ID")
	clientSecret := utils.GetEnvVar("TWITCH_CLIENT_SECRET")
	return clientId, clientSecret
}

func ListenToMessages(conn *websocket.Conn, wsMessages chan struct{}) {
	defer close(wsMessages)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}

		log.Println("Received message:", string(message))

		var msgData map[string]interface{}
		err = json.Unmarshal(message, &msgData)
		if err != nil {
			log.Fatal("JSON parsing error:", err)
		}

		// Check if the message is a welcome message
		if metadata, ok := msgData["metadata"].(map[string]interface{}); ok {
			if messageType, ok := metadata["message_type"].(string); ok && messageType == twitchWelcomeMessage {
				HandleWelcomeMessage(conn)
			}
		}
	}
}

func HandleWelcomeMessage(conn *websocket.Conn) {
	log.Println("Sending response to welcome message")

	// Get Oauth token if needed and begin to subscribe to events
	SendOauthRequest()
}

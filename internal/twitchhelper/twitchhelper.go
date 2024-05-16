// Main helper to connect to Twitch events
package twitchhelper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

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
	err, oauthResponse := SendOauthRequest()

	if err != nil {
		log.Println("Error sending Oauth request:", err)
		return
	}

	log.Println("Oauth response:", oauthResponse)
}

// GetLiveChannelsCount fetches the current number of live channels from Twitch API
func GetLiveChannelsCount(clientId string, oauthToken string) (int, error) {

	u := url.URL{Scheme: "https", Host: "api.twitch.tv", Path: "/helix/streams"}
	q := u.Query()
	q.Set("type", "live")
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Authorization", "Bearer "+oauthToken)

	// user_name is the Twitch username of the broadcaster
	// user_id is the id that you might want?

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %v", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling JSON response: %v", err)
	}

	fmt.Printf("Data is %v\n", data)

	return 0, nil
}

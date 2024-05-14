// Main helper to connect to Twitch events
package twitchhelper

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/michaelfioretti/twitch-stats-producer/internal/utils"
)

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
		log.Printf("Received message: %s", message)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("HELLOOOOOO HERE WE GO!")
	secretPath := "/run/secrets/db_password"
	secrets, err := os.ReadFile(secretPath)
	if err != nil {
		log.Fatalf("Failed to read secret: %v", err)
	}

	dbPasswordStr := string(secrets)
	dbPasswordStr = strings.TrimSpace(dbPasswordStr)

	fmt.Println("FINALLY: ", dbPasswordStr)

	// go kafkahelper.ValidateBaseTopics()

	// shared.TwitchClient = twitchchatparser.CreateTwitchClient()
	// twitchchatparser.SubscribeToTwitchChat()

	// shared.TwitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
	// 	twitchMessage := twitchchatparser.ParseTwitchMessage(message)
	// 	shared.MessageChannel <- twitchMessage
	// })

	// go twitchchatparser.ProcessTwitchMessages()
	// go shared.TwitchClient.Connect()

	// http.ListenAndServe(":8080", nil)
}

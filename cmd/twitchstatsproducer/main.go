package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
)

func main() {
	go kafkahelper.ValidateBaseTopics()

	// First, get the OAuth token
	err, oauthToken := twitchhelper.SendOauthRequest()

	if err != nil {
		log.Fatal("Error getting OAuth token:", err)
	}

	token := "oauth:" + oauthToken.AccessToken

	// Fetch the top 100 streamers, and begin parsing their Twitch chat
	streamerChatDataChan := make(chan models.IRCChatMessageData)
	doneChan := make(chan struct{})

	topLivestreams, err := twitchhelper.GetTop100Livestreams(oauthToken.AccessToken)
	if err != nil {
		log.Fatalf("Error getting livestreams: %v\n", err)
	}

	fmt.Print("Here are the top 100 livestreams, the game, and the streamer\n")
	fmt.Print("Count: ", len(topLivestreams), "\n\n")
	for i, stream := range topLivestreams {
		fmt.Printf("%d: Streamer: %s, with: %d viewers\n", i, stream.UserName, stream.ViewerCount)
	}

	for _, stream := range topLivestreams {
		streamer := stream.UserName
		conn, err := net.Dial("tcp", constants.TWITCH_IRC_URL)
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}

		// 4. Authenticate and join channels
		fmt.Fprintf(conn, "PASS %s\r\n", token)
		fmt.Fprintf(conn, "NICK %s\r\n", constants.TWITCH_USERNAME) // Your Twitch username
		fmt.Fprintf(conn, "JOIN #%s\r\n", strings.ToLower(streamer))

		go twitchhelper.ReadStreamerChat(streamer, conn, streamerChatDataChan)
		go twitchhelper.ProcessStreamerChat(streamerChatDataChan)
	}

	<-doneChan
}

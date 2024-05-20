package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
)

func readStreamerChat(streamer string, conn net.Conn, streamerMsgChannel chan<- models.IRCChatMessageData) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// Check if the line is a PING message (keep the connection alive)
		if strings.HasPrefix(line, "PING") {
			fmt.Fprintf(conn, "PONG :tmi.twitch.tv\r\n")
		} else {
			// Process the chat message here (e.g., print it)
			fmt.Printf("[%s] %s", streamer, line)
			streamerMsgChannel <- models.IRCChatMessageData{Streamer: streamer, Message: line}
		}
	}
}

// processData receives data from the channel and processes it.
func processData(dataChan <-chan models.IRCChatMessageData) {
	fmt.Println("Processing data...")
	for data := range dataChan {
		fmt.Printf("Channel: %s,  Message: %s\n", data.Streamer, data.Message)
		// Your data processing logic here
	}
}

func main() {
	go kafkahelper.ValidateBaseTopics()

	// First, get the OAuth token
	err, oauthToken := twitchhelper.SendOauthRequest()

	if err != nil {
		log.Fatal("Error getting OAuth token:", err)
	}

	token := "oauth:" + oauthToken.AccessToken

	// Fetch the top 100 streamers, and begin parsing their Twitch chat
	dataChan := make(chan models.IRCChatMessageData)
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

		go readStreamerChat(streamer, conn, dataChan)
		go processData(dataChan)
	}

	<-doneChan
}

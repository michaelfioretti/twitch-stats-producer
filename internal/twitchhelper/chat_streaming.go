package twitchhelper

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
)

func ReadStreamerChat(streamer string, conn net.Conn, streamerMsgChannel chan<- models.IRCChatMessageData) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Print(line)
		// Check if the line is a PING message (keep the connection alive)
		if strings.HasPrefix(line, "PING") {
			fmt.Fprintf(conn, "%s\r\n", constants.TWITCH_PONG_URL)
		} else {
			streamerMsgChannel <- models.IRCChatMessageData{Streamer: streamer, Message: line}
		}
	}
}

// processData receives data from the channel and processes it.
func ProcessStreamerChat(dataChan <-chan models.IRCChatMessageData) {
	for data := range dataChan {
		fmt.Printf("Channel: %s,  Message: %s\n", data.Streamer, data.Message)
		// Your data processing logic here
	}
}

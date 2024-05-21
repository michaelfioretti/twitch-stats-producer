package twitchhelper

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
)

func ReadStreamerChat(streamer string, conn net.Conn, streamerMsgChannel chan<- models.IRCChatMessageData) {
	fmt.Printf("Now reading streamer %s chat\n", streamer)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print("Got a line?????")
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Print(line)
		// Check if the line is a PING message (keep the connection alive)
		if strings.HasPrefix(line, "PING") {
			fmt.Fprintf(conn, "%s\r\n", constants.TWITCH_PONG_URL)
		} else {
			streamerMsgChannel <- models.IRCChatMessageData{Streamer: streamer, Message: line, Timestamp: time.Now().String()}
		}
	}
}

// processData receives data from the channel and processes it.
func ProcessStreamerChat(dataChan <-chan models.IRCChatMessageData) {
	fmt.Print("need to process...")
	for data := range dataChan {
		// Your data processing logic here
		fmt.Sprintf("Channel: %s,  Message: %s, Timestamp: %s\n", data.Streamer, data.Message, data.Timestamp)
		// msgStr := fmt.Sprintf("Channel: %s,  Message: %s, Timestamp: %s\n", data.Streamer, data.Message, data.Timestamp)
		// msg := kafka.Message{Value: []byte(msgStr)}
		// kafkaproducer.WriteDataToKafka("streamer_chat", []kafka.Message{msg})
	}
}

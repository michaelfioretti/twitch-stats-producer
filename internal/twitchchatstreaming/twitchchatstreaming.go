package twitchchatstreaming

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
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
			streamerMsgChannel <- models.IRCChatMessageData{Streamer: streamer, Message: line, Timestamp: time.Now().String()}
		}
	}
}

/*
ProcessStreamerChat will take the base message, timestamp, etc and format it
into a ParsedTwitchMessage struct. This struct will then be sent to the
appropriate Kafka topic.
*/
func ProcessStreamerChat(dataChan <-chan models.IRCChatMessageData) {
	for data := range dataChan {
		// First, check to see if we should process the message. If yes, then
		// we will parse and send it to the appropriate Kafka topic.
		if twitchchatparser.ShouldProcessMessage(data.Message) {
			parsedMessage := twitchchatparser.ParseMessage(data.Message)
			fmt.Printf("%v", parsedMessage)
			fmt.Sprintf("Channel: %s,  Message: %s, Timestamp: %s\n", data.Streamer, data.Message, data.Timestamp)
			// msgStr := fmt.Sprintf("Channel: %s,  Message: %s, Timestamp: %s\n", data.Streamer, data.Message, data.Timestamp)
			// msg := kafka.Message{Value: []byte(msgStr)}
			// kafkaproducer.WriteDataToKafka("streamer_chat", []kafka.Message{msg})
		}
	}
}

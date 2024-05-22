package twitchchatstreaming

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkaproducer"
	"github.com/michaelfioretti/twitch-stats-producer/internal/models"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
	"github.com/segmentio/kafka-go"
)

func ReadStreamerChat(streamer string, conn net.Conn, streamerMsgChannel chan<- models.IRCChatMessageData) {
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

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
			parsedChatMessage := twitchchatparser.ParseMessage(data.Message)

			// Now, we need to format the message into a string that can be sent to Kafka
			chatMsg, err := json.Marshal(parsedChatMessage)
			if err != nil {
				fmt.Println("Error marshalling message:", err)
				return
			}

			msg := kafka.Message{Value: []byte(chatMsg)}
			kafkaproducer.WriteDataToKafka("streamer_chat", []kafka.Message{msg})
		}
	}
}

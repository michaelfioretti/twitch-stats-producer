package twitchchatparser

import (
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	kafkahelper "github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/shared"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func CreateTwitchClient() *twitch.Client {
	// First, get the OAuth token
	oauthToken := twitchhelper.SendOauthRequest()

	token := "oauth:" + oauthToken.AccessToken

	client := twitch.NewClient(constants.TWITCH_USERNAME, token)
	client.Capabilities = []string{twitch.TagsCapability, twitch.CommandsCapability, twitch.MembershipCapability} // Customize which capabilities are sent

	return client
}

// Used to get the initial top 100 streamers. Subsequent updates will be done in the UpdateStreamerList method
func SubscribeToTwitchChat() {
	topLivestreams := twitchhelper.GetTop100ChannelsByStreamViewCount()
	go shared.TwitchClient.Join(topLivestreams...)
}

func ProcessTwitchMessages() {
	go func() {
		for msg := range shared.MessageChannel {
			twitchMessage, err := proto.Marshal(msg)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Error marshaling message")

				continue
			}

			msg := kafka.Message{Value: twitchMessage}
			shared.KafkaMessageBatch = append(shared.KafkaMessageBatch, msg)

			if len(shared.KafkaMessageBatch) == constants.KAFKA_MESSAGES_PER_BATCH {
				log.Info("Writing 100 more messages at this time: ", time.Now().Format("2006-01-02 15:04:05"))

				go kafkahelper.WriteDataToKafka("streamer_chat", shared.KafkaMessageBatch)

				shared.ProcessedMessageCount += constants.KAFKA_MESSAGES_PER_BATCH

				if shared.ProcessedMessageCount >= constants.TWITCH_RESET_STREAM_MESSAGE_COUNT {
					go UpdateStreamerList(shared.TwitchClient)
					shared.ProcessedMessageCount = 0
				}

				shared.KafkaMessageBatch = make([]kafka.Message, 0)
			}
		}
	}()
}

// We will go through each of the streams that we are currently watching and update the list of top 100 streams
// to watch. This will be done every 50,000 messages that we process.
func UpdateStreamerList(client *twitch.Client) {
	shared.LastUpdatedTopStreamersMutex.Lock()
	defer shared.LastUpdatedTopStreamersMutex.Unlock()

	topChannelsByViewCount := twitchhelper.GetTop100ChannelsByStreamViewCount()

	// Get the difference between the two lists
	var streamsToJoin, streamsToLeave []string

	// Use a map to store the current connected channels for faster lookup
	connectedChannelsMap := make(map[string]bool)
	topChannelsByViewCountMap := make(map[string]bool)
	for _, channel := range shared.LastUpdatedTopStreamers {
		connectedChannelsMap[channel] = true
	}

	for _, channel := range topChannelsByViewCount {
		topChannelsByViewCountMap[channel] = true
	}

	// Find the channels to leave and join
	for _, channel := range topChannelsByViewCount {
		if _, found := connectedChannelsMap[channel]; !found {
			streamsToJoin = append(streamsToJoin, channel)
		}
	}

	for _, channel := range shared.LastUpdatedTopStreamers {
		if _, found := topChannelsByViewCountMap[channel]; !found {
			streamsToLeave = append(streamsToLeave, channel)
		}
	}

	// Leave the channels that are no longer in the top 100
	if len(streamsToJoin) > 0 {
		shared.TwitchClient.Join(streamsToJoin...)
	}

	for _, channel := range streamsToLeave {
		shared.TwitchClient.Depart(channel)
	}

	shared.LastUpdatedTopStreamers = topChannelsByViewCount
}

func ParseTwitchMessage(message twitch.PrivateMessage) *models.TwitchMessage {
	// Extract the username from the message sender.
	username := message.User.Name

	// Get the message content.
	messageText := message.Message

	// Check if the message is an action (/me).
	isAction := strings.HasPrefix(messageText, "\x01ACTION ")
	if isAction {
		// Remove the action prefix to get the actual message.
		messageText = strings.TrimPrefix(messageText, "\x01ACTION ")
		messageText = strings.TrimSuffix(messageText, "\x01")
	}

	// Extract badges from the message tags.
	var badges []string
	badgeStr := message.Tags["badges"]
	if badgeStr != "" {
		badges = strings.Split(badgeStr, ",")
	}

	// Extract the number of bits cheered (if any).
	bitsCheered, _ := strconv.ParseInt(message.Tags["bits"], 10, 32)

	// Get the channel name from the message tags.
	channel := message.Tags["room-id"]

	subscribed, _ := strconv.Atoi(message.Tags["subscriber"])
	mod, _ := strconv.Atoi(message.Tags["mod"])

	// Return the structured data.
	return &models.TwitchMessage{
		Username:   username,
		Channel:    message.Channel,
		Message:    messageText,
		Badges:     badges,
		Bits:       int32(bitsCheered),
		Mod:        int32(mod),
		Subscribed: int32(subscribed),
		Color:      message.Tags["color"],
		RoomID:     "#" + channel,
	}
}

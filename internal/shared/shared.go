package shared

import (
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/segmentio/kafka-go"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
)

var MessageChannel = make(chan *models.TwitchMessage, constants.TWITCH_MESSAGE_CHANNEL_BUFFER_SIZE)
var KafkaMessageBatch = make([]kafka.Message, 0)
var ProcessedMessageCount = 0
var TwitchClient *twitch.Client

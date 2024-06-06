package shared

import (
	"sync"

	"github.com/segmentio/kafka-go"

	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
)

var MessageChannel = make(chan *models.TwitchMessage, 1000)
var KafkaMessageBatch = make([]kafka.Message, 0)
var KafkaMessageBatchMutex = sync.Mutex{}

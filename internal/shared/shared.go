package shared

import (
	"sync"

	firebase "firebase.google.com/go"
	"github.com/gempir/go-twitch-irc/v2"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
)

var MessageChannel = make(chan *models.TwitchMessage, constants.MESSAGES_PER_BATCH)
var TwitchClient *twitch.Client
var LastUpdatedTopStreamers []string
var LastUpdatedTopStreamersMutex *sync.RWMutex = &sync.RWMutex{}
var FirebaseConfigPath string = "twitch-chat-stats-firebase-adminsdk-zm2i4-23ecdda6eb.json"
var FirebaseApp *firebase.App

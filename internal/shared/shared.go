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
var FirebaseConfigPath string = "twitch-chats-firebase-adminsdk-fbsvc-d5c364858b.json"
var FirebaseApp *firebase.App

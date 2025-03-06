package firebasehelper

import (
	"context"
	"sync"
	"time"

	firebase "firebase.google.com/go"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/shared"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type FirebaseMessageBatcher struct {
	mu          sync.Mutex
	messages    []*models.TwitchMessage
	client      *firebase.App
	ctx         context.Context
	maxMessages int
}

var messageBatcher FirebaseMessageBatcher

func ConnectToFirebase() {
	opt := option.WithCredentialsFile(shared.FirebaseConfigPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	shared.FirebaseApp = app

	messageBatcher = FirebaseMessageBatcher{
		client:   app,
		ctx:      context.Background(),
		messages: make([]*models.TwitchMessage, 0, constants.MESSAGES_PER_BATCH),
	}

	go messageBatchFlusher()
}

func saveMessageBatchToFirebase() {
	log.Infof("Saving %d messages to firebase!!!", len(messageBatcher.messages))
	ctx := context.Background()
	client, err := shared.FirebaseApp.Firestore(ctx)

	if err != nil {
		log.Fatalf("error getting Firestore client: %v", err)
	}

	defer client.Close()

	batch := client.BulkWriter(ctx)

	for _, msg := range messageBatcher.messages {
		ref := client.Collection("stream_chat").NewDoc()
		ref.Set(ctx, msg)
		batch.Create(ref, models.TwitchMessage{})
	}

	log.Info("Now commiting...")

	batch.End()
}

func messageBatchFlusher() {
	for {
		time.Sleep(constants.FLUSH_INTERVAL)
		messageBatcher.mu.Lock()

		if len(messageBatcher.messages) > 0 {
			// saveMessageBatchToFirebase()
		}

		messageBatcher.messages = make([]*models.TwitchMessage, 0, messageBatcher.maxMessages)
		messageBatcher.mu.Unlock()
	}
}

func ProcessTwitchMessages() {
	for msg := range shared.MessageChannel {
		messageBatcher.messages = append(messageBatcher.messages, msg)

		// if len(messageBatcher.messages) >= constants.MESSAGES_PER_BATCH {
		// 	log.Infof("Writing %d more messages at this time: ", constants.MESSAGES_PER_BATCH, time.Now().Format("2006-01-02 15:04:05"))

		// 	saveMessageBatchToFirebase()
		// }
	}
}

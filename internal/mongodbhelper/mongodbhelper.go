package mongodbhelper

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/shared"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbMessageBatcher struct {
	mu          sync.Mutex
	messages    []*models.TwitchMessage
	db          *mongo.Client
	ctx         context.Context
	maxMessages int
}

var messageBatcher MongoDbMessageBatcher

func ConnectToMongoDb() {
	DB_NAME, DB_PASSWORD := loadDatabaseKeys()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.z9zf2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", DB_NAME, DB_PASSWORD)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	// Ensure the client is connected
	err = db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Infof("Connected to MongoDB")

	messageBatcher = MongoDbMessageBatcher{
		db:       db,
		ctx:      context.Background(),
		messages: make([]*models.TwitchMessage, 0, constants.MESSAGES_PER_BATCH),
	}

	go messageBatchFlusher()
}

func DisconnectFromMongoDn() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := messageBatcher.db.Disconnect(ctx)
	if err != nil {
		log.Fatalf("Failed to disconnect from MongoDB: %v", err)
	}
	fmt.Println("Disconnected from MongoDB")
}

func saveMessageBatchToMongoDb() {
	log.Infof("Saving %d messages to Mongo: ", len(messageBatcher.messages))

	coll := messageBatcher.db.Database("twitch_chat_stats").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newMessages := []interface{}{}
	for _, msg := range messageBatcher.messages {
		// Convert Protobuf timestamp to time.Time for MongoDB
		doc := bson.M{
			"username":   msg.Username,
			"channel":    msg.Channel,
			"message":    msg.Message,
			"badges":     msg.Badges,
			"bits":       msg.Bits,
			"mod":        msg.Mod,
			"subscribed": msg.Subscribed,
			"color":      msg.Color,
			"roomID":     msg.RoomID,
			"createdat":  msg.CreatedAt.AsTime(),
		}
		newMessages = append(newMessages, doc)
	}

	_, err := coll.InsertMany(ctx, newMessages)

	if err != nil {
		log.Fatal(err)
	}
}

func messageBatchFlusher() {
	for {
		time.Sleep(constants.FLUSH_INTERVAL)
		messageBatcher.mu.Lock()

		if len(messageBatcher.messages) > 0 {
			saveMessageBatchToMongoDb()
		}

		// Update total processed messages and reset streamers to watch
		shared.TotalMessageCount += int(len(messageBatcher.messages))
		log.Infof("Total messages processed: %d out of limit of %d", shared.TotalMessageCount, constants.TWITCH_RESET_STREAM_MESSAGE_COUNT)
		if shared.TotalMessageCount >= constants.TWITCH_RESET_STREAM_MESSAGE_COUNT {
			twitchchatparser.UpdateStreamerList(shared.TwitchClient)
			shared.TotalMessageCount = 0
		}

		// Reset message batch
		messageBatcher.messages = make([]*models.TwitchMessage, 0, messageBatcher.maxMessages)
		messageBatcher.mu.Unlock()
	}
}

func ProcessTwitchMessages() {
	for msg := range shared.MessageChannel {
		messageBatcher.mu.Lock()
		messageBatcher.messages = append(messageBatcher.messages, msg)
		messageBatcher.mu.Unlock()
	}
}

func loadDatabaseKeys() (string, string) {
	// Note: in production, env variables will be injected in
	if err := godotenv.Load(); err != nil {
		logrus.Debug("No .env file found, continuing...")
	}

	return os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
}

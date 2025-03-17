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

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://michaelfioretti1:McUzouT2shcg2wLJ@cluster0.z9zf2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	db, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := db.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	log.Infof("Pinged your deployment. You successfully connected to MongoDB!")

	messageBatcher = MongoDbMessageBatcher{
		db:       db,
		ctx:      context.Background(),
		messages: make([]*models.TwitchMessage, 0, constants.MESSAGES_PER_BATCH),
	}

	go messageBatchFlusher()
}

func saveMessageBatchToMongoDb() {
	log.Infof("Saving %d messages to Mongo: ", len(messageBatcher.messages))

	DB_NAME, DB_PASSWORD := loadDatabaseKeys()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.z9zf2.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", DB_NAME, DB_PASSWORD)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	db, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	if err := db.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	coll := db.Database("twitch_chat_stats").Collection("messages")

	newMessages := []interface{}{}
	for _, msg := range messageBatcher.messages {
		newMessages = append(newMessages, msg)
	}

	_, err = coll.InsertMany(context.TODO(), newMessages)

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

		messageBatcher.messages = make([]*models.TwitchMessage, 0, messageBatcher.maxMessages)
		messageBatcher.mu.Unlock()
	}
}

func ProcessTwitchMessages() {
	for msg := range shared.MessageChannel {
		log.Print(msg)
		messageBatcher.mu.Lock()
		messageBatcher.messages = append(messageBatcher.messages, msg)
		messageBatcher.mu.Unlock()
	}
}

func loadDatabaseKeys() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD")
}

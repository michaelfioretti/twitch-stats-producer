package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/michaelfioretti/twitch-stats-producer/internal/kafkahelper"
	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchchatparser"
)

func main() {
	go kafkahelper.ValidateBaseTopics()

	client := twitchchatparser.CreateTwitchClient()

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")

	// We will update the top 100 streamers every 5 minutes
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Updating top 100 streamers at this time: ", time.Now().Format("2006-01-02 15:04:05"))
				twitchchatparser.UpdateStreamerList(client)
			}
		}
	}()
}

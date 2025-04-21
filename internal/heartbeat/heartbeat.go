package heartbeat

import (
	"log"
	"net/http"
	"time"
)

func StartHeartbeat(url string) {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Error sending heartbeat: %v", err)
				continue
			}
			resp.Body.Close()
			log.Printf("Heartbeat sent to %s, status: %s", url, resp.Status)
		}
	}()
}

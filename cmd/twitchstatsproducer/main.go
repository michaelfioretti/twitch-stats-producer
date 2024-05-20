package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/michaelfioretti/twitch-stats-producer/internal/twitchhelper"
)

func main() {
	// go kafkahelper.ValidateBaseTopics()
	// Helper code to keep server running (for now!)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// First, get the OAuth token
	err, oauthToken := twitchhelper.SendOauthRequest()

	if err != nil {
		log.Fatal("Error getting OAuth token:", err)
	}

	// 1. Get your OAuth token (replace with your actual token)
	token := "oauth:" + oauthToken.AccessToken

	// 2. List of streamers to connect to
	streamers := []string{"piratesoftware"}

	for _, streamer := range streamers {
		// 3. Connect to the Twitch IRC server
		conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}
		defer conn.Close()

		// 4. Authenticate and join channels
		fmt.Fprintf(conn, "PASS %s\r\n", token)
		fmt.Fprintf(conn, "NICK justinfan12345\r\n") // Your Twitch username
		fmt.Fprintf(conn, "JOIN #%s\r\n", strings.ToLower(streamer))

		// 5. Read and process chat messages
		go func(streamer string, conn net.Conn) {
			reader := bufio.NewReader(conn)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading:", err)
					return
				}

				// Check if the line is a PING message (keep the connection alive)
				if strings.HasPrefix(line, "PING") {
					fmt.Fprintf(conn, "PONG :tmi.twitch.tv\r\n")
				} else {
					// Process the chat message here (e.g., print it)
					fmt.Printf("[%s] %s", streamer, line)
				}
			}
		}(streamer, conn)
	}

	for {
		// Keep the program running indefinitely
	}

	// topLivestreams, err := twitchhelper.GetTop100Livestreams(oauthToken.AccessToken)
	// if err != nil {
	// 	fmt.Printf("Error getting livestreams: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Print("Here are the top 100 livestreams, the game, and the streamer\n")
	// fmt.Print("Count: ", len(topLivestreams), "\n\n")
	// for i, stream := range topLivestreams {
	// 	fmt.Printf("%d: Streamer: %s, with: %d viewers\n", i, stream.UserName, stream.ViewerCount)
	// }

	// u := url.URL{Scheme: "wss", Host: constants.TWITCH_PUBSUB_URL}
	// headers := http.Header{}
	// headers.Add("Authorization", "Bearer "+oauthToken.AccessToken)

	// log.Printf("Securely connecting to %s", u.String())

	// conn, _, err := websocket.DefaultDialer.Dial(u.String(), headers)
	// if err != nil {
	// 	log.Fatal("Dial error:", err)
	// }

	// defer conn.Close()

	// // Subscribe to a specific topic
	// subscribeMessage := `{"type": "LISTEN", "data": {"topics": ["` + topic + `"], "auth_token": "` + accessToken + `"}}`
	// err = conn.WriteMessage(websocket.TextMessage, []byte(subscribeMessage))
	// if err != nil {
	// 	return fmt.Errorf("writing subscribe message failed: %v", err)
	// }

	// fmt.Println("Connected to Twitch PubSub API. Waiting for messages...")

	// // Read messages from the connection
	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		return fmt.Errorf("error reading message: %v", err)
	// 	}
	// 	fmt.Printf("Received message: %s\n", message)
	// 	// Process the received message here (e.g., parse JSON and handle accordingly)
	// }

	// wsMessages := make(chan struct{})

	// go twitchhelper.ListenToMessages(conn, wsMessages)

	// // Channel management based on messages received from Twitch
	// for {
	// 	select {
	// 	case <-wsMessages:
	// 		return
	// 	case <-interrupt:
	// 		log.Println("Interrupt received, closing connection...")
	// 		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	// 		if err != nil {
	// 			log.Println("Write close error:", err)
	// 		}
	// 		select {
	// 		case <-wsMessages:
	// 		}
	// 		return
	// 	}
	// }
}

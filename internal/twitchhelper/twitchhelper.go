// Main helper to connect to Twitch events
package twitchhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	"github.com/michaelfioretti/twitch-stats-producer/internal/utils"

	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	log "github.com/sirupsen/logrus"
)

func SendOauthRequest() *models.TwitchOauthResponse {
	var oauthResponse models.TwitchOauthResponse
	clientId, clientSecret := loadTwitchKeys()
	req, err := http.NewRequest("POST", constants.TWITCH_OAUTH_URL, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	q.Add("client_id", clientId)
	q.Add("client_secret", clientSecret)
	q.Add("grant_type", constants.TWITCH_OAUTH_REQUEST_TYPE)

	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &oauthResponse)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	return &oauthResponse
}

func GetTop100ChannelsByStreamViewCount() []string {
	oauthToken := SendOauthRequest()

	var top100Streams models.Top100StreamsResponse
	clientId, _ := loadTwitchKeys()

	u := url.URL{Scheme: "https", Host: "api.twitch.tv", Path: "/helix/streams"}
	q := u.Query()
	q.Add("first", "100")
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Authorization", "Bearer "+oauthToken.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &top100Streams)
	if err != nil {
		log.Fatalf("Error unmarshaling response body: %v", err)
	}

	streamerNames := make([]string, 0, 100)
	totalViewCount := int32(0)

	for _, stream := range top100Streams.Data {
		totalViewCount = totalViewCount + stream.ViewerCount
		streamerNames = append(streamerNames, stream.UserName)
		log.Info("Streamer name: ", stream.UserName, " - View count: ", stream.ViewerCount)
	}

	log.Info("Total view count for top 100 streams: ", totalViewCount)

	return streamerNames
}

func loadTwitchKeys() (string, string) {
	secrets := utils.ReadSecrets()
	return secrets["twitch_client_id"], secrets["twitch_client_secret"]
}

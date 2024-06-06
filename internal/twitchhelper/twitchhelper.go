// Main helper to connect to Twitch events
package twitchhelper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
	models "github.com/michaelfioretti/twitch-stats-producer/internal/models/proto"
	"github.com/michaelfioretti/twitch-stats-producer/internal/utils"
)

const (
	twitchWelcomeMessage = "session_welcome"
)

func SendOauthRequest() *models.TwitchOauthResponse {
	var oauthResponse models.TwitchOauthResponse
	clientId, clientSecret := LoadTwitchKeys()
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

func LoadTwitchKeys() (string, string) {
	// Load Twitch keys from .env file
	clientId := utils.GetEnvVar("TWITCH_CLIENT_ID")
	clientSecret := utils.GetEnvVar("TWITCH_CLIENT_SECRET")
	return clientId, clientSecret
}

func GetTop100Livestreams(oauthToken string) (*models.Top100StreamsResponse, error) {
	if oauthToken == "" {
		return nil, fmt.Errorf("OAuth token is required")
	}

	var top100Streams models.Top100StreamsResponse
	clientId, _ := LoadTwitchKeys()

	u := url.URL{Scheme: "https", Host: "api.twitch.tv", Path: "/helix/streams"}
	q := u.Query()
	q.Set("first", "100")
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Authorization", "Bearer "+oauthToken)

	resp, err := client.Do(req)
	if err != nil {
		return &top100Streams, err
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

	return &top100Streams, nil
}

func GetTrendingGames(accessToken string) ([]models.TwitchGame, error) {
	clientId, _ := LoadTwitchKeys()
	// Twitch API Endpoint for Get Top Games
	url := "https://api.twitch.tv/helix/games/top"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set Twitch API Client ID (Get yours from the Twitch Developer Console)
	req.Header.Set("Client-Id", clientId)

	// Set Authorization Header if you have a Twitch App Access Token (Optional)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Twitch API Error: %s", resp.Status)
	}

	// Parse the JSON response
	var response struct {
		Data []models.TwitchGame `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

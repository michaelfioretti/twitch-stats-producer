package twitchhelper

import (
	"encoding/json"
	"net/http"

	"github.com/michaelfioretti/twitch-stats-producer/internal/constants"
)

type TwitchOauthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func SendOauthRequest() (error, TwitchOauthResponse) {
	clientId, clientSecret := LoadTwitchKeys()
	req, err := http.NewRequest("POST", constants.TWITCH_OAUTH_URL, nil)
	if err != nil {
		return err, TwitchOauthResponse{}
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
		return err, TwitchOauthResponse{}
	}

	defer resp.Body.Close()

	var oauthResponse TwitchOauthResponse
	err = json.NewDecoder(resp.Body).Decode(&oauthResponse)
	if err != nil {
		return err, TwitchOauthResponse{}
	}

	return nil, oauthResponse
}

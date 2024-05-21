package twitchhelper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSendOauthRequest tests the SendOauthRequest function
func TestSendOauthRequest(t *testing.T) {
	clientId, clientSecret := LoadTwitchKeys()
	if clientId == "" || clientSecret == "" {
		t.Skip("Twitch keys not set")
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "some_token", "expires_in": 12345, "token_type": "bearer"}`))
	}))
	defer server.Close()

	err, oauthResponse := SendOauthRequest()
	assert.Nil(t, err)
	assert.Equal(t, "some_token", oauthResponse.AccessToken)
}

// TestGetTop100Livestreams tests the GetTop100Livestreams function
func TestGetTop100Livestreams(t *testing.T) {
	oauthToken := "some_token"
	streams, err := GetTop100Livestreams(oauthToken)
	assert.Nil(t, err)
	assert.NotNil(t, streams)
}

// TestGetTrendingGames tests the GetTrendingGames function
func TestGetTrendingGames(t *testing.T) {
	accessToken := "some_token"
	games, err := GetTrendingGames(accessToken)
	assert.Nil(t, err)
	assert.NotNil(t, games)
}

// TestLoadTwitchKeys tests the LoadTwitchKeys function
func TestLoadTwitchKeys(t *testing.T) {
	clientId, clientSecret := LoadTwitchKeys()
	assert.NotEmpty(t, clientId)
	assert.NotEmpty(t, clientSecret)
}

// TestTwitchMessageRequestMarshaling tests the marshaling of TwitchMessageRequest
func TestTwitchMessageRequestMarshaling(t *testing.T) {
	request := TwitchMessageRequest{
		Type:    twitchWelcomeMessage,
		Version: "1.0",
		Condition: struct {
			BroadcasterUserId string `json:"broadcaster_user_id"`
			ModeratorUserId   string `json:"moderator_user_id"`
		}{
			BroadcasterUserId: "123456789",
			ModeratorUserId:   "987654321",
		},
	}

	marshaled, err := json.Marshal(request)
	assert.Nil(t, err)

	unmarshaled := TwitchMessageRequest{}
	err = json.Unmarshal(marshaled, &unmarshaled)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(request, unmarshaled))
}

// TestTwitchOauthResponseMarshaling tests the marshaling of TwitchOauthResponse
func TestTwitchOauthResponseMarshaling(t *testing.T) {
	response := TwitchOauthResponse{
		AccessToken: "some_token",
		ExpiresIn:   12345,
		TokenType:   "bearer",
	}

	marshaled, err := json.Marshal(response)
	assert.Nil(t, err)

	unmarshaled := TwitchOauthResponse{}
	err = json.Unmarshal(marshaled, &unmarshaled)
	assert.Nil(t, err)

	assert.True(t, reflect.DeepEqual(response, unmarshaled))
}

// TestGetTop100LivestreamsError tests the GetTop100Livestreams function with an error
func TestGetTop100LivestreamsError(t *testing.T) {
	oauthToken := ""
	streams, err := GetTop100Livestreams(oauthToken)
	assert.NotNil(t, err)
	assert.Nil(t, streams)
}

// TestGetTrendingGamesError tests the GetTrendingGames function with an error
func TestGetTrendingGamesError(t *testing.T) {
	accessToken := ""
	games, err := GetTrendingGames(accessToken)
	assert.NotNil(t, err)
	assert.Nil(t, games)
}

// TestSendOauthRequestError tests the SendOauthRequest function with an error
func TestSendOauthRequestError(t *testing.T) {
	clientId, clientSecret := LoadTwitchKeys()
	if clientId == "" || clientSecret == "" {
		t.Skip("Twitch keys not set")
	}

	err, oauthResponse := SendOauthRequest()
	assert.NotNil(t, err)
	assert.Empty(t, oauthResponse.AccessToken)
}

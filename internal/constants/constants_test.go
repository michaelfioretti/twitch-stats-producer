package constants

import "testing"

func TestTwitchOauthURL(t *testing.T) {
	expected := "https://id.twitch.tv/oauth2/token"
	if TwitchOauthURL != expected {
		t.Errorf("Expected TwitchOauthURL to be %s, but got %s", expected, TwitchOauthURL)
	}
}

func TestTwitchEventSubWSS(t *testing.T) {
	expected := "eventsub.wss.twitch.tv"
	if TwitchEventSubWSS != expected {
		t.Errorf("Expected TwitchEventSubWSS to be %s, but got %s", expected, TwitchEventSubWSS)
	}
}

func TestTwitchOauthRequestType(t *testing.T) {
	expected := "client_credentials"
	if TwitchOauthRequestType != expected {
		t.Errorf("Expected TwitchOauthRequestType to be %s, but got %s", expected, TwitchOauthRequestType)
	}
}

func TestTwitchUsersAPIUrl(t *testing.T) {
	expected := "https://api.twitch.tv/helix/users"
	if TwitchUsersAPIUrl != expected {
		t.Errorf("Expected TwitchUsersAPIUrl to be %s, but got %s", expected, TwitchUsersAPIUrl)
	}
}

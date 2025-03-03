package constants

import (
	"testing"
	"time"
)

func TestTwitchOauthUrl(t *testing.T) {
	expected := "https://id.twitch.tv/oauth2/token"
	if TWITCH_OAUTH_URL != expected {
		t.Errorf("Expected TWITCH_OAUTH_URL to be %s, but got %s", expected, TWITCH_OAUTH_URL)
	}
}

func TestTwitchOauthRequestType(t *testing.T) {
	expected := "client_credentials"
	if TWITCH_OAUTH_REQUEST_TYPE != expected {
		t.Errorf("Expected TWITCH_OAUTH_REQUEST_TYPE to be %s, but got %s", expected, TWITCH_OAUTH_REQUEST_TYPE)
	}
}

func TestTwitchIrcUrl(t *testing.T) {
	expected := "irc.chat.twitch.tv:6667"
	if TWITCH_IRC_URL != expected {
		t.Errorf("Expected TWITCH_IRC_URL to be %s, but got %s", expected, TWITCH_IRC_URL)
	}
}

func TestTwitchUsername(t *testing.T) {
	expected := "justinfan12345"
	if TWITCH_USERNAME != expected {
		t.Errorf("Expected TWITCH_USERNAME to be %s, but got %s", expected, TWITCH_USERNAME)
	}
}

func TestTwitchTagsRequestCmd(t *testing.T) {
	expected := "CAP REQ :twitch.tv/commands twitch.tv/tags"
	if TWITCH_TAGS_REQUEST_CMD != expected {
		t.Errorf("Expected TWITCH_TAGS_REQUEST_CMD to be %s, but got %s", expected, TWITCH_TAGS_REQUEST_CMD)
	}
}

func TestTwitchResetStreamMessageCount(t *testing.T) {
	expected := 50000
	if TWITCH_RESET_STREAM_MESSAGE_COUNT != expected {
		t.Errorf("Expected TWITCH_RESET_STREAM_MESSAGE_COUNT to be %d, but got %d", expected, TWITCH_RESET_STREAM_MESSAGE_COUNT)
	}
}

func TestFlushInterval(t *testing.T) {
	expected := 5 * time.Second
	if FLUSH_INTERVAL != expected {
		t.Errorf("Expected FLUSH_INTERVAL to be %s, but got %s", expected, FLUSH_INTERVAL)
	}
}

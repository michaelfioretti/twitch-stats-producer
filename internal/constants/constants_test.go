package constants

import "testing"

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

func TestKafkaTopics(t *testing.T) {
	expected := "top_streamers,streamer_stats,streamer_chat,viewer_demographics,trending_games"
	if KAFKA_TOPICS != expected {
		t.Errorf("Expected KAFKA_TOPICS to be %s, but got %s", expected, KAFKA_TOPICS)
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

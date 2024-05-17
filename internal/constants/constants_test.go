package constants

import "testing"

func TestNumber1(t *testing.T) {
	expected := "1"
	if NUMBER_1_STR != expected {
		t.Errorf("Expected NUMBER_1_STR to be %s, but got %s", expected, NUMBER_1_STR)
	}
}

func TestWebsocketString(t *testing.T) {
	expected := "websocket"
	if WEBSOCKET_STRING != expected {
		t.Errorf("Expected WEBSOCKET_STRING to be %s, but got %s", expected, WEBSOCKET_STRING)
	}

}
func TestTwitchOauthUrl(t *testing.T) {
	expected := "https://id.twitch.tv/oauth2/token"
	if TWITCH_OAUTH_URL != expected {
		t.Errorf("Expected TWITCH_OAUTH_URL to be %s, but got %s", expected, TWITCH_OAUTH_URL)
	}
}

func TestTwitchEventSubWss(t *testing.T) {
	expected := "eventsub.wss.twitch.tv"
	if TWITCH_EVENT_SUB_WSS != expected {
		t.Errorf("Expected TWITCH_EVENT_SUB_WSS to be %s, but got %s", expected, TWITCH_EVENT_SUB_WSS)
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

func TestTwitchUsersApiUrl(t *testing.T) {
	expected := "https://api.twitch.tv/helix/users"
	if TWITCH_USERS_API_URL != expected {
		t.Errorf("Expected TWITCH_USERS_API_URL to be %s, but got %s", expected, TWITCH_USERS_API_URL)
	}
}

func TestTwitchChannelChatSubscriptionType(t *testing.T) {
	expected := "channel.chat.message"
	if TWITCH_CHANNEL_CHAT_SUBSCRIPTION_TYPE != expected {
		t.Errorf("Expected TWITCH_CHANNEL_CHAT_SUBSCRIPTION_TYPE to be %s, but got %s", expected, TWITCH_CHANNEL_CHAT_SUBSCRIPTION_TYPE)
	}
}

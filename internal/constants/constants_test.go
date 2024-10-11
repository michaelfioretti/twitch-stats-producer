package constants

import "testing"

func TestKafkaTopics(t *testing.T) {
	expected := "streamer_chat"
	if KAFKA_TOPICS != expected {
		t.Errorf("Expected KAFKA_TOPICS to be %s, but got %s", expected, KAFKA_TOPICS)
	}
}

func TestKafkaMessagesPerBatch(t *testing.T) {
	expected := 100
	if MESSAGES_PER_BATCH != expected {
		t.Errorf("Expected KAFKA_MESSAES_PER_BATCH to be %d, but got %d", expected, MESSAGES_PER_BATCH)
	}
}

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

func TestTwitchMessageChannelBufferSize(t *testing.T) {
	expected := 1000
	if TWITCH_MESSAGE_CHANNEL_BUFFER_SIZE != expected {
		t.Errorf("Expected TWITCH_MESSAGE_CHANNEL_BUFFER_SIZE to be %d, but got %d", expected, TWITCH_MESSAGE_CHANNEL_BUFFER_SIZE)
	}
}
func TestTwitchResetStreamMessageCount(t *testing.T) {
	expected := 50000
	if TWITCH_RESET_STREAM_MESSAGE_COUNT != expected {
		t.Errorf("Expected TWITCH_RESET_STREAM_MESSAGE_COUNT to be %d, but got %d", expected, TWITCH_RESET_STREAM_MESSAGE_COUNT)
	}
}

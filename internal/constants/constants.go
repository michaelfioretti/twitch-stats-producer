package constants

const (
	KAFKA_TOPICS                       = "streamer_chat"
	KAFKA_MESSAGES_PER_BATCH           = 100
	PARTITION_COUNT                    = 1
	REPLICATION_COUNT                  = 1
	TWITCH_OAUTH_URL                   = "https://id.twitch.tv/oauth2/token"
	TWITCH_OAUTH_REQUEST_TYPE          = "client_credentials"
	TWITCH_IRC_URL                     = "irc.chat.twitch.tv:6667"
	TWITCH_USERNAME                    = "justinfan12345"
	TWITCH_TAGS_REQUEST_CMD            = "CAP REQ :twitch.tv/commands twitch.tv/tags"
	TWITCH_MESSAGE_CHANNEL_BUFFER_SIZE = 1000
	TWITCH_RESET_STREAM_MESSAGE_COUNT  = 50000
)

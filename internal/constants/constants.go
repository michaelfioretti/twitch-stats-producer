package constants

const (
	// Topics list - used when creating/initializing Kafka topics for persistent/real-time data storage
	KAFKA_TOPICS = "top_streamers,streamer_stats,streamer_chat,viewer_demographics,trending_games"
	// Twitch
	TWITCH_OAUTH_URL          = "https://id.twitch.tv/oauth2/token"
	TWITCH_OAUTH_REQUEST_TYPE = "client_credentials"
	TWITCH_IRC_URL            = "irc.chat.twitch.tv:6667"
	TWITCH_USERNAME           = "justinfan12345"
	TWITCH_PONG_URL           = "PONG :tmi.twitch.tv"
	TWITCH_TAGS_REQUEST_CMD   = "CAP REQ :twitch.tv/commands twitch.tv/tags"
	// Number of messages that we will produce before updating the top 100 streams to watch
	TWITCH_RESET_STREAM_MESSAGE_COUNT = 50000
)

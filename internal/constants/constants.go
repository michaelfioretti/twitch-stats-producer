package constants

const (
	// Magic strings
	WEBSOCKET_STRING = "websocket"
	NUMBER_1_STR     = "1"
	// Twitch
	TWITCH_OAUTH_URL          = "https://id.twitch.tv/oauth2/token"
	TWITCH_USERS_API_URL      = "https://api.twitch.tv/helix/users"
	TWITCH_PUBSUB_URL         = "pubsub-edge.twitch.tv"
	TWITCH_EVENT_SUB_WSS      = "eventsub.wss.twitch.tv"
	TWITCH_OAUTH_REQUEST_TYPE = "client_credentials"
	// Topics list - used when creating/initializing Kafka topics for persistent/real-time data storage
	KAFKA_TOPICS = "top_streamers,streamer_stats,streamer_chat,viewer_demographics,trending_games"
	// Specific subscription types
	TWITCH_CHANNEL_CHAT_SUBSCRIPTION_TYPE = "channel.chat.message"
	TWITCH_IRC_URL                        = "irc.chat.twitch.tv:6667"
	TWITCH_USERNAME                       = "justinfan12345"
)

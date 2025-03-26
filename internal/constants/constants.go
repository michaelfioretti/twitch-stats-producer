package constants

import "time"

const (
	TWITCH_OAUTH_URL                  = "https://id.twitch.tv/oauth2/token"
	TWITCH_OAUTH_REQUEST_TYPE         = "client_credentials"
	TWITCH_IRC_URL                    = "irc.chat.twitch.tv:6667"
	TWITCH_USERNAME                   = "justinfan12345"
	TWITCH_TAGS_REQUEST_CMD           = "CAP REQ :twitch.tv/commands twitch.tv/tags"
	MESSAGES_PER_BATCH                = 500
	TWITCH_RESET_STREAM_MESSAGE_COUNT = 500000
	FLUSH_INTERVAL                    = 5 * time.Second
)

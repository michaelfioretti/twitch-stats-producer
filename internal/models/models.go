package models

// Data is a struct to hold parsed IRC data.
type IRCChatMessageData struct {
	// Add fields as needed (e.g., Command, Channel, User, Message)
	Streamer  string
	Message   string
	Timestamp string
}

type Top100StreamsResponse struct {
	Data []Stream `json:"data"`
}

type Stream struct {
	ID           string   `json:"id"`
	UserID       string   `json:"user_id"`
	UserName     string   `json:"user_name"`
	GameID       string   `json:"game_id"`
	Type         string   `json:"type"`
	Title        string   `json:"title"`
	ViewerCount  int      `json:"viewer_count"`
	StartedAt    string   `json:"started_at"`
	Language     string   `json:"language"`
	ThumbnailURL string   `json:"thumbnail_url"`
	TagIDs       []string `json:"tag_ids"`
	IsMature     bool     `json:"is_mature"`
}

// TwitchGame represents a game object from the Twitch API
type TwitchGame struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BoxArtURL   string `json:"box_art_url"`
	ViewerCount int    `json:"viewer_count"`
	Popularity  int    `json:"popularity"`
}

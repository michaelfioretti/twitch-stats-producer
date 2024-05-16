package models

type Stream struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
	GameID      string `json:"game_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	ViewerCount int    `json:"viewer_count"`
	StartedAt   string `json:"started_at"`
	Language    string `json:"language"`
	Thumbnail   struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"thumbnail_url"`
}

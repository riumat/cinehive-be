package dto

type ExternalIdsDto struct {
	InstagramID string `json:"instagram_id,omitempty"`
	TwitterID   string `json:"twitter_id,omitempty"`
	FacebookID  string `json:"facebook_id,omitempty"`
	YoutubeID   string `json:"youtube_id,omitempty"`
	TikTokID    string `json:"tiktok_id,omitempty"`
}

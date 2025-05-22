package types

type CastItem struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
}

type CastItemWithCharacter struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
	Character   string  `json:"character"`
}

type CrewItem struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
	Job         string  `json:"job"`
}

type GenericPersonItem struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
}
type EpisodeItem struct {
	Name    string  `json:"name"`
	Id      float64 `json:"id"`
	AirDate string  `json:"air_date"`
}

type SeasonItem struct {
	Id           float64 `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	EpisodeCount int     `json:"episode_count"`
	AirDate      string  `json:"air_date"`
}

type LanguageItem struct {
	ID          string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name,omitempty"`
}

type ExternalIdsItem struct {
	InstagramID string `json:"instagram_id,omitempty"`
	TwitterID   string `json:"twitter_id,omitempty"`
	FacebookID  string `json:"facebook_id,omitempty"`
	YoutubeID   string `json:"youtube_id,omitempty"`
	TikTokID    string `json:"tiktok_id,omitempty"`
}

type RecommendedItem struct {
	ID           float64 `json:"id"`
	Title        string  `json:"title,omitempty"`
	Name         string  `json:"name,omitempty"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	GenreIds     []int   `json:"genre_ids"`
	MediaType    string  `json:"media_type"`
}

type VideoItem struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type ProviderItem struct {
	LogoPath     string  `json:"logo_path"`
	ProviderName string  `json:"provider_name"`
	ProviderID   float64 `json:"provider_id"`
}

type ProductionCompanyItem struct {
	Name     string  `json:"name"`
	LogoPath string  `json:"logo_path"`
	ID       float64 `json:"id"`
}

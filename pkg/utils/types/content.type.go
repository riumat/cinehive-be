package types

type Person struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
}

type Actor struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
	Character   string  `json:"character"`
}

type CrewMember struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
	Job         string  `json:"job"`
}

type Episode struct {
	Name    string  `json:"name"`
	Id      float64 `json:"id"`
	AirDate string  `json:"air_date"`
}

type Season struct {
	Id           float64 `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	EpisodeCount int     `json:"episode_count"`
	AirDate      string  `json:"air_date"`
}

type LanguageSpoken struct {
	ID          string `json:"iso_639_1"`
	Name        string `json:"name"`
	EnglishName string `json:"english_name,omitempty"`
}

type Video struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	PublishedAt string `json:"published_at"`
	Size        int    `json:"size"`
	Type        string `json:"type"`
}

type Provider struct {
	LogoPath     string  `json:"logo_path"`
	ProviderName string  `json:"provider_name"`
	ProviderID   float64 `json:"provider_id"`
}

type ProductionCompany struct {
	Name     string  `json:"name"`
	LogoPath string  `json:"logo_path"`
	ID       float64 `json:"id"`
}

type ProductionCountry struct {
	Name string `json:"name"`
	Code string `json:"iso_3166_1"`
}

type ContentSummary struct {
	ID           float64 `json:"id"`
	Title        string  `json:"title,omitempty"`
	Name         string  `json:"name,omitempty"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	GenreIds     []int   `json:"genre_ids"`
	MediaType    string  `json:"media_type"`
}

type ContentCard struct {
	ID           int     `json:"id"`
	Title        string  `json:"title,omitempty"`
	Name         string  `json:"name,omitempty"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	VoteAverage  float64 `json:"vote_average"`
	ReleaseDate  string  `json:"release_date,omitempty"`
	FirstAirDate string  `json:"first_air_date,omitempty"`
	MediaType    string  `json:"media_type"`
}

type SearchResult struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title,omitempty"`
	Name               string  `json:"name,omitempty"`
	PosterPath         string  `json:"poster_path"`
	VoteAverage        float64 `json:"vote_average,omitempty"`
	ReleaseDate        string  `json:"release_date,omitempty"`
	FirstAirDate       string  `json:"first_air_date,omitempty"`
	MediaType          string  `json:"media_type"`
	ProfilePath        string  `json:"profile_path,omitempty"`
	KnownForDepartment string  `json:"known_for_department,omitempty"`
}

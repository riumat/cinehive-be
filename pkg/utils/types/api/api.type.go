package api

import "github.com/riumat/cinehive-be/pkg/utils/types"

type DetailsResponse struct {
	Id               float64                   `json:"id"`
	Name             string                    `json:"name,omitempty"`
	Title            string                    `json:"title,omitempty"`
	Overview         string                    `json:"overview"`
	ReleaseDate      string                    `json:"release_date,omitempty"`
	FirstAirDate     string                    `json:"first_air_date,omitempty"`
	Genres           []types.Genre             `json:"genres"`
	BackdropPath     string                    `json:"backdrop_path"`
	PosterPath       string                    `json:"poster_path"`
	NextEpisodeToAir *types.EpisodeItem        `json:"next_episode_to_air,omitempty"`
	LastEpisodeToAir *types.EpisodeItem        `json:"last_episode_to_air,omitempty"`
	Seasons          *[]types.SeasonItem       `json:"seasons,omitempty"`
	SpokenLanguages  []types.LanguageItem      `json:"spoken_languages"`
	Status           string                    `json:"status"`
	Runtime          float64                   `json:"runtime,omitempty"`
	Budget           float64                   `json:"budget,omitempty"`
	Revenue          float64                   `json:"revenue,omitempty"`
	CreatedBy        []types.GenericPersonItem `json:"created_by,omitempty"`
	Credits          struct {
		Cast []types.CastItem `json:"cast"`
		Crew []types.CrewItem `json:"crew"`
	} `json:"credits"`
	ExternalIds         types.ExternalIdsItem         `json:"external_ids"`
	ProductionCompanies []types.ProductionCompanyItem `json:"production_companies"`
	ProductionCountries []struct {
		Name string `json:"name"`
		Code string `json:"iso_3166_1"`
	} `json:"production_countries"`
	Providers       []types.ProviderItem    `json:"providers"`
	Recommendations []types.RecommendedItem `json:"recommendations"`
	Videos          []types.VideoItem       `json:"videos"`
	VoteAverage     float64                 `json:"vote_average"`
	VoteCount       float64                 `json:"vote_count"`
}

type CastResponse struct {
	Id           float64                       `json:"id"`
	Name         string                        `json:"name,omitempty"`
	Title        string                        `json:"title,omitempty"`
	BackdropPath string                        `json:"backdrop_path"`
	PosterPath   string                        `json:"poster_path"`
	Cast         []types.CastItemWithCharacter `json:"cast"`
}

type CrewResponse struct {
	Id           float64          `json:"id"`
	Name         string           `json:"name,omitempty"`
	Title        string           `json:"title,omitempty"`
	BackdropPath string           `json:"backdrop_path"`
	PosterPath   string           `json:"poster_path"`
	Crew         []types.CrewItem `json:"crew"`
}

type VideoResponse struct {
	Id           float64 `json:"id"`
	Name         string  `json:"name,omitempty"`
	Title        string  `json:"title,omitempty"`
	BackdropPath string  `json:"backdrop_path"`
	PosterPath   string  `json:"poster_path"`
	Videos       struct {
		Trailers []types.VideoItem `json:"trailers"`
		Others   []types.VideoItem `json:"others"`
	} `json:"videos"`
}

type RecommendationsResponse struct {
	Id              float64                 `json:"id"`
	Name            string                  `json:"name,omitempty"`
	Title           string                  `json:"title,omitempty"`
	BackdropPath    string                  `json:"backdrop_path"`
	PosterPath      string                  `json:"poster_path"`
	Recommendations []types.RecommendedItem `json:"recommendations"`
}

type SeasonResponse struct {
	Id           float64            `json:"id"`
	Name         string             `json:"name"`
	BackdropPath string             `json:"backdrop_path"`
	PosterPath   string             `json:"poster_path"`
	Seasons      []types.SeasonItem `json:"seasons"`
}

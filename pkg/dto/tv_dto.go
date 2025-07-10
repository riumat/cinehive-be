package dto

import "github.com/riumat/cinehive-be/pkg/utils/types"

type TvDetailsDto struct {
	Id                  float64                   `json:"id"`
	Name                string                    `json:"name"`
	Overview            string                    `json:"overview"`
	FirstAirDate        string                    `json:"first_air_date"`
	Genres              []types.Genre             `json:"genres"`
	BackdropPath        string                    `json:"backdrop_path"`
	PosterPath          string                    `json:"poster_path"`
	NextEpisodeToAir    types.Episode             `json:"next_episode_to_air"`
	LastEpisodeToAir    types.Episode             `json:"last_episode_to_air"`
	Seasons             []types.Season            `json:"seasons"`
	SpokenLanguages     []types.LanguageSpoken    `json:"spoken_languages"`
	Status              string                    `json:"status"`
	CreatedBy           []types.Person            `json:"created_by"`
	Credits             CreditsDto                `json:"credits"`
	ExternalIds         ExternalIdsDto            `json:"external_ids"`
	ProductionCompanies []types.ProductionCompany `json:"production_companies"`
	ProductionCountries []types.ProductionCountry `json:"production_countries"`
	Recommendations     []types.ContentSummary    `json:"recommendations"`
	Videos              []types.Video             `json:"videos"`
	VoteAverage         float64                   `json:"vote_average"`
	VoteCount           float64                   `json:"vote_count"`
	ProvidersLink       string                    `json:"providers_link"`
	MediaType           string                    `json:"media_type,omitempty"`
}

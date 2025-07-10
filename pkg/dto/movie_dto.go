package dto

import "github.com/riumat/cinehive-be/pkg/utils/types"

type MovieDetailsDto struct {
	Id                  float64                   `json:"id"`
	Title               string                    `json:"title"`
	Overview            string                    `json:"overview"`
	ReleaseDate         string                    `json:"release_date"`
	Genres              []types.Genre             `json:"genres"`
	BackdropPath        string                    `json:"backdrop_path"`
	PosterPath          string                    `json:"poster_path"`
	SpokenLanguages     []types.LanguageSpoken    `json:"spoken_languages"`
	Status              string                    `json:"status"`
	Runtime             float64                   `json:"runtime"`
	Budget              float64                   `json:"budget"`
	Revenue             float64                   `json:"revenue"`
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

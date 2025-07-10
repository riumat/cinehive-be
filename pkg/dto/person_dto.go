package dto

import "github.com/riumat/cinehive-be/pkg/utils/types"

type PersonDto struct {
	ID                 int64                 `json:"id"`
	Name               string                `json:"name"`
	ProfilePath        string                `json:"profile_path"`
	KnownForDepartment string                `json:"known_for_department"`
	Birthday           string                `json:"birthday"`
	Deathday           string                `json:"deathday"`
	PlaceOfBirth       string                `json:"place_of_birth"`
	CastCredits        []ContentPersonWorkTo `json:"cast_credits"`
	CrewCredits        []ContentPersonWorkTo `json:"crew_credits"`
	KnowFor            []types.ContentCard   `json:"known_for"`
	ExternalIds        ExternalIdsDto        `json:"external_ids"`
	HeaderBackdrop     ContentPersonWorkTo   `json:"header_backdrop"`
}

type ContentPersonWorkTo struct {
	ID           float64 `json:"id"`
	Title        string  `json:"title,omitempty"`
	Name         string  `json:"name,omitempty"`
	MediaType    string  `json:"media_type"`
	Character    string  `json:"character,omitempty"`
	Job          string  `json:"job,omitempty"`
	ReleaseDate  string  `json:"release_date,omitempty"`
	VoteAverage  float64 `json:"vote_average"`
	BackdropPath string  `json:"backdrop_path"`
	VoteCount    float64 `json:"vote_count"`
}

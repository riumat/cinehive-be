package dto

import "github.com/riumat/cinehive-be/pkg/utils/types"

type CreditsDto struct {
	Cast []types.Actor      `json:"cast"`
	Crew []types.CrewMember `json:"crew"`
}

type HomeListsDto struct {
	TrendingMovies []types.ContentCard `json:"trending_movies"`
	TrendingTv     []types.ContentCard `json:"trending_tv"`
	TopRatedMovies []types.ContentCard `json:"top_rated_movies"`
	TopRatedTv     []types.ContentCard `json:"top_rated_tv"`
}

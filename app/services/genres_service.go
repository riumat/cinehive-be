package services

import (
	"fmt"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type CustomResponse struct {
	Genres []Genre `json:"genres"`
}

func FetchGenres(client *config.TMDBClient, mediaType string) ([]Genre, error) {
	resp, err := HttpGet[CustomResponse](client, endpoints.TmdbEndpoint.Genre.All(mediaType), nil)

	if err != nil {
		return []Genre{}, fmt.Errorf("failed to fetch data from TMDB: %w", err)
	}

	return resp.Genres, nil
}

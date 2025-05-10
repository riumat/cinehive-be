package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FetchGenres(client *config.TMDBClient, mediaType string) ([]Genre, error) {
	resp, err := client.Get(endpoints.TmdbEndpoint.Genre.All(mediaType), nil)

	if err != nil {
		return []Genre{}, fmt.Errorf("failed to fetch data from TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code %d", resp.StatusCode)
	}

	var result struct {
		Genres []Genre `json:"genres"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []Genre{}, fmt.Errorf("failed to parse TMDB response: %w", err)
	}

	return result.Genres, nil
}

package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
)

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	GenresID    []int  `json:"genre_ids"`
}

func FetchFeaturedMovie(client *config.TMDBClient) (Movie, error) {
	resp, err := client.Get(endpoints.TmdbEndpoint.Trending.Movies, nil)
	if err != nil {
		return Movie{}, fmt.Errorf("failed to fetch data from TMDB: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Movie{}, fmt.Errorf("TMDB API returned status code %d", resp.StatusCode)
	}

	var result struct {
		Results []Movie `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Movie{}, fmt.Errorf("failed to parse TMDB response: %w", err)
	}

	return result.Results[0], nil
}

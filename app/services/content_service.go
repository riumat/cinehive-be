package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"golang.org/x/sync/errgroup"
)

type ContentCard struct {
	ID         int    `json:"id"`
	Title      string `json:"title,omitempty"`
	Name       string `json:"name,omitempty"`
	PosterPath string `json:"poster_path"`
}

func ContentFetcher(client *config.TMDBClient, url string) ([]ContentCard, error) {
	resp, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("TMDB API returned status code %d", resp.StatusCode)
	}

	var result struct {
		Results []ContentCard `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

func FetchLandingCards(client *config.TMDBClient) (map[string][]ContentCard, error) {
	var g errgroup.Group
	var movies, tvShows []ContentCard

	g.Go(func() error {
		results, err := ContentFetcher(client, endpoints.TmdbEndpoint.Trending.Movies)
		if err != nil {
			return err
		}
		movies = results
		return nil
	})

	g.Go(func() error {
		results, err := ContentFetcher(client, endpoints.TmdbEndpoint.Trending.TV)
		if err != nil {
			return err
		}
		tvShows = results
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return map[string][]ContentCard{
		"movies": movies,
		"tv":     tvShows,
	}, nil
}

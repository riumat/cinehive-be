package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type ContentCard struct {
	ID         int    `json:"id"`
	Title      string `json:"title,omitempty"`
	Name       string `json:"name,omitempty"`
	PosterPath string `json:"poster_path"`
}

func FetchLandingCards(client *config.TMDBClient) (map[string][]ContentCard, error) {
	var g errgroup.Group
	var movies, tvShows []ContentCard

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.Trending.Movies, nil)
		if err != nil {
			return err
		}
		movies = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.Trending.TV, nil)
		if err != nil {
			return err
		}
		tvShows = results.Results
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

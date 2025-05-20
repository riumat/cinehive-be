package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type ContentCard struct {
	ID           int     `json:"id"`
	Title        string  `json:"title,omitempty"`
	Name         string  `json:"name,omitempty"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	VoteAverage  float64 `json:"vote_average"`
	ReleaseDate  string  `json:"release_date,omitempty"`
	FirstAirDate string  `json:"first_air_date,omitempty"`
	MediaType    string  `json:"media_type"`
}

type Provider struct {
	IT struct {
		Flatrate []struct {
			LogoPath        string  `json:"logo_path"`
			ProviderID      float64 `json:"provider_id"`
			ProviderName    string  `json:"provider_name"`
			ProviderCountry string  `json:"provider_country"`
		} `json:"flatrate"`
	}
}

func FetchLandingCards(client *config.TMDBClient) (map[string][]ContentCard, error) {
	var g errgroup.Group
	var trendingMovies, trendingTvShows, topRatedMovies, topRatedTvShows []ContentCard

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.Trending.Movies, nil)
		if err != nil {
			return err
		}
		for i := range results.Results {
			results.Results[i].MediaType = "movie"
		}
		trendingMovies = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.Trending.TV, nil)
		if err != nil {
			return err
		}
		for i := range results.Results {
			results.Results[i].MediaType = "tv"
		}
		trendingTvShows = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.TopRated.Movies, nil)
		if err != nil {
			return err
		}
		for i := range results.Results {
			results.Results[i].MediaType = "movie"
		}
		topRatedMovies = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]ContentCard]](client, endpoints.TmdbEndpoint.TopRated.TV, nil)
		if err != nil {
			return err
		}
		for i := range results.Results {
			results.Results[i].MediaType = "tv"
		}
		topRatedTvShows = results.Results
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return map[string][]ContentCard{
		"trendingMovies": trendingMovies,
		"trendingTv":     trendingTvShows,
		"topRatedMovies": topRatedMovies,
		"topRatedTv":     topRatedTvShows,
	}, nil
}

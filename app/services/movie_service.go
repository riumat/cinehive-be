package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"golang.org/x/sync/errgroup"
)

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	GenresID    []int  `json:"genre_ids"`
}

type Provider struct {
	IT struct {
		Flatrate []struct {
			LogoPath        string `json:"logo_path"`
			ProviderID      int    `json:"provider_id"`
			ProviderName    string `json:"provider_name"`
			ProviderCountry string `json:"provider_country"`
		} `json:"flatrate"`
	}
}

func FetchFeaturedMovie(client *config.TMDBClient) (any, error) {

	data, err := HttpGet[utils.Response[[]Movie]](client, endpoints.TmdbEndpoint.Trending.Movies, nil)
	if err != nil {
		return nil, err
	}

	return data.Results[0], nil
}

func FetchMovieHeaderDetails(client *config.TMDBClient, id string) (any, error) {
	var g errgroup.Group
	var details, images map[string]any
	var providers Provider

	providerParams := map[string]string{
		"watch_region": "IT",
	}

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend("movie", id, []string{"external_ids"}), nil)
		if err != nil {
			return err
		}
		details = results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[Provider]](client, endpoints.TmdbEndpoint.DynamicContent.Providers("movie", id), providerParams)
		if err != nil {
			return err
		}
		providers = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Images("movie", id), nil)
		if err != nil {
			return err
		}
		images = results
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	details["providers"] = providers.IT.Flatrate
	details["images"] = images

	return details, nil

}

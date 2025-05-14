package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/helpers"
	"github.com/riumat/cinehive-be/pkg/utils/types"
	"golang.org/x/sync/errgroup"
)

type Tv struct {
	Title       string `json:"name"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	GenresID    []int  `json:"genre_ids"`
}

const (
	TV = "tv"
)

func FetchTvHeaderDetails(client *config.TMDBClient, id string) (any, error) {
	var g errgroup.Group
	var details, images map[string]any
	var providers Provider

	providerParams := map[string]string{
		"watch_region": "IT",
	}

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(TV, id, []string{"external_ids"}), nil)
		if err != nil {
			return err
		}
		details = results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[Provider]](client, endpoints.TmdbEndpoint.DynamicContent.Providers(TV, id), providerParams)
		if err != nil {
			return err
		}
		providers = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Images(TV, id), nil)
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

func FetchTvOverviewDetails(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(TV, id, []string{"credits"}), nil)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func FetchTvCastDetails(client *config.TMDBClient, id string) ([]types.CastItem, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Credits(TV, id, "aggregate_credits"), nil)
	if err != nil {
		return nil, err
	}

	cast, ok := results["cast"].([]any)
	if !ok {
		return nil, nil
	}

	return helpers.ExtractCastItems(cast), nil
}

func FetchTvCrewDetails(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Credits(TV, id, "aggregate_credits"), nil)
	if err != nil {
		return nil, err
	}

	crew, ok := results["crew"].([]any)
	if !ok {
		return crew, nil
	}

	formattedCrew := utils.FormatCrewTvList(crew)

	return formattedCrew, nil
}

func FetchTvVideos(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Videos(TV, id), nil)
	if err != nil {
		return nil, err
	}

	videos, ok := results["results"].([]any)
	if !ok {
		return nil, nil
	}

	type MovieVideos struct {
		Trailers []any `json:"trailers"`
		Others   []any `json:"others"`
	}

	trailers, others := utils.FormatVideoList(videos)

	movieVideos := MovieVideos{
		Trailers: trailers,
		Others:   others,
	}

	return movieVideos, nil
}

func FetchTvRecommendations(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[utils.Response[[]Tv]](client, endpoints.TmdbEndpoint.DynamicContent.Recommendations(TV, id), nil)
	if err != nil {
		return nil, err
	}

	return results.Results, nil
}

func FetchTvSeasons(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.All(TV, id), nil)
	if err != nil {
		return nil, err
	}

	return results["seasons"], nil
}

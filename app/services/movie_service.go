package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
	"golang.org/x/sync/errgroup"
)

type Movie struct {
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
	PosterPath  string `json:"poster_path"`
	GenresID    []int  `json:"genre_ids"`
}

const (
	MOVIE = "movie"
)

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
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"external_ids"}), nil)
		if err != nil {
			return err
		}
		details = results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[utils.Response[Provider]](client, endpoints.TmdbEndpoint.DynamicContent.Providers(MOVIE, id), providerParams)
		if err != nil {
			return err
		}
		providers = results.Results
		return nil
	})

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Images(MOVIE, id), nil)
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

func FetchMovieOverviewDetails(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"credits"}), nil)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func FetchMovieCastDetails(client *config.TMDBClient, id string) ([]types.CastItem, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Credits(MOVIE, id, "credits"), nil)
	if err != nil {
		return nil, err
	}
	cast, ok := results["cast"].([]any)
	if !ok {
		return nil, nil
	}

	var castItems []types.CastItem
	for _, item := range cast {
		if castMap, ok := item.(map[string]interface{}); ok {
			castItem := types.CastItem{}
			if id, ok := castMap["id"].(float64); ok {
				castItem.ID = id
			}
			if name, ok := castMap["name"].(string); ok {
				castItem.Name = name
			}
			if profile, ok := castMap["profile_path"].(string); ok {
				castItem.ProfilePath = profile
			}
			if character, ok := castMap["character"].(string); ok {
				castItem.Character = character
			}
			castItems = append(castItems, castItem)
		}
	}

	return castItems, nil
}

func FetchMovieCrewDetails(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Credits(MOVIE, id, "credits"), nil)
	if err != nil {
		return nil, err
	}

	crew, ok := results["crew"].([]any)
	if !ok {
		return crew, nil
	}

	formattedCrew := utils.FormatMovieCrewList(crew)

	return formattedCrew, nil
}

func FetchMovieVideos(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.Videos(MOVIE, id), nil)
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

func FetchMovieRecommendations(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[utils.Response[[]Movie]](client, endpoints.TmdbEndpoint.DynamicContent.Recommendations(MOVIE, id), nil)
	if err != nil {
		return nil, err
	}

	return results.Results, nil
}

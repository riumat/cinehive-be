package services

import (
	"encoding/json"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
	"github.com/riumat/cinehive-be/pkg/utils/types/api"
	"golang.org/x/sync/errgroup"
)

type Movie struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	ReleaseDate  string `json:"release_date"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path,omitempty"`
	GenresID     []int  `json:"genre_ids"`
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

func FetchMovieDetails(client *config.TMDBClient, id string) (api.DetailsResponse, error) {
	var g errgroup.Group
	var details map[string]any
	var providers Provider

	providerParams := map[string]string{
		"watch_region": "IT",
	}

	g.Go(func() error {
		results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"external_ids", "credits", "recommendations", "videos"}), nil)
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

	if err := g.Wait(); err != nil {
		return api.DetailsResponse{}, err
	}

	details["providers"] = providers.IT.Flatrate

	if recs, ok := details["recommendations"].(map[string]any); ok {
		details["recommendations"] = recs["results"]
	}
	if videos, ok := details["videos"].(map[string]any); ok {
		trailers, _ := utils.FormatVideoList(videos["results"].([]any))
		details["videos"] = trailers
	}

	jsonBytes, err := json.Marshal(details)
	if err != nil {
		return api.DetailsResponse{}, err
	}

	var resp api.DetailsResponse
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return api.DetailsResponse{}, err
	}

	return resp, nil
}

func FetchMovieCastDetails(client *config.TMDBClient, id string) (api.CastResponse, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"credits"}), nil)

	if err != nil {
		return api.CastResponse{}, err
	}
	if credits, ok := results["credits"].(map[string]any); ok {
		results["cast"] = credits["cast"]
	}

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return api.CastResponse{}, err
	}

	var resp api.CastResponse
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return api.CastResponse{}, err
	}

	return resp, nil
}

func FetchMovieCrewDetails(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"credits"}), nil)
	if err != nil {
		return api.CrewResponse{}, err
	}

	var crewItems []types.CrewItem
	if credits, ok := results["credits"].(map[string]any); ok {
		if crewArr, ok := credits["crew"].([]any); ok {
			crewItems = utils.FormatMovieCrewList(crewArr)
		}
	}

	results["crew"] = crewItems

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return api.CrewResponse{}, err
	}

	var resp api.CrewResponse
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return api.CrewResponse{}, err
	}

	return resp, nil
}

func FetchMovieVideos(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"videos"}), nil)
	if err != nil {
		return api.VideoResponse{}, err
	}

	if videos, ok := results["videos"].(map[string]any); ok {
		if res, ok := videos["results"].([]any); ok {
			results["videos"] = res
		}
	}

	type MovieVideos struct {
		Trailers []any `json:"trailers"`
		Others   []any `json:"others"`
	}

	trailers, others := utils.FormatVideoList(results["videos"].([]any))

	movieVideos := MovieVideos{
		Trailers: trailers,
		Others:   others,
	}

	results["videos"] = movieVideos

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return api.VideoResponse{}, err
	}

	var resp api.VideoResponse
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return api.VideoResponse{}, err
	}

	return resp, nil
}

func FetchMovieRecommendations(client *config.TMDBClient, id string) (any, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"recommendations"}), nil)
	if err != nil {
		return api.RecommendationsResponse{}, err
	}

	if recs, ok := results["recommendations"].(map[string]any); ok {
		if res, ok := recs["results"].([]any); ok {
			results["recommendations"] = res
		}
	}

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return api.RecommendationsResponse{}, err
	}

	var resp api.RecommendationsResponse
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return api.RecommendationsResponse{}, err
	}

	return resp, nil
}

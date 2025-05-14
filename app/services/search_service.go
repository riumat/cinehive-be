package services

import (
	"fmt"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

type SearchResult struct {
	Results      []any `json:"results"`
	TotalPages   int   `json:"total_pages"`
	TotalResults int   `json:"total_results"`
	Page         int   `json:"page"`
}

func FetchSearchResults(client *config.TMDBClient, query, page string) (utils.PaginatedResponse[[]any], error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := HttpGet[utils.PaginatedResponse[[]any]](client, endpoints.TmdbEndpoint.Search.Multi, queryParams)
	if err != nil {
		return utils.PaginatedResponse[[]any]{}, fmt.Errorf("failed to fetch data from TMDB: %w", err)
	}

	return resp, nil
}

func FetchSearchWithFilters(client *config.TMDBClient, params types.FilterParams, media string) (utils.PaginatedResponse[[]any], error) {
	var release string
	if media == "movie" {
		release = "release_date"
	} else if media == "tv" {
		release = "first_air_date"
	} else {
		return utils.PaginatedResponse[[]any]{}, fmt.Errorf("invalid media type")
	}

	queryParams := map[string]string{
		"page":                         params.Page,
		"with_watch_providers":         params.Providers,
		"with_genres":                  params.Genres,
		fmt.Sprintf("%s.gte", release): fmt.Sprintf("%s-01-01", params.From),
		fmt.Sprintf("%s.lte", release): fmt.Sprintf("%s-12-31", params.To),
		"sort_by":                      params.Sort,
		"without_genres":               "10763,10764,10767",
		"vote_count.gte":               "200",
		"with_runtime.gte":             params.RuntimeGte,
		"with_runtime.lte":             params.RuntimeLte,
		"watch_region":                 "IT",
	}

	resp, err := HttpGet[utils.PaginatedResponse[[]any]](client, endpoints.TmdbEndpoint.Discover.All(media), queryParams)
	if err != nil {
		return utils.PaginatedResponse[[]any]{}, fmt.Errorf("failed to fetch filtered content from TMDB: %w", err)
	}

	return resp, nil
}

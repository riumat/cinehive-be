package services

import (
	"fmt"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
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

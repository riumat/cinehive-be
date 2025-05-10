package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
)

type SearchResult struct {
	Results      []any `json:"results"`
	TotalPages   int   `json:"total_pages"`
	TotalResults int   `json:"total_results"`
	Page         int   `json:"page"`
}

func FetchSearchResults(client *config.TMDBClient, query, page string) (SearchResult, error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := client.Get(endpoints.TmdbEndpoint.Search.Multi, queryParams)
	if err != nil {
		return SearchResult{}, fmt.Errorf("failed to fetch data from TMDB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SearchResult{}, fmt.Errorf("TMDB API returned status code %d", resp.StatusCode)
	}

	var result struct {
		Results      []any `json:"results"`
		TotalPages   int   `json:"total_pages"`
		TotalResults int   `json:"total_results"`
		Page         int   `json:"page"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return SearchResult{}, fmt.Errorf("failed to parse TMDB response: %w", err)
	}

	return result, nil
}

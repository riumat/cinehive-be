package services

import (
	"fmt"
	"log"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

func FetchSearchResults(client *config.TMDBClient, query, page string) (dto.SearchDto, error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := HttpGet[dto.SearchDto](client, endpoints.TmdbEndpoint.Search.Multi, queryParams)
	if err != nil {
		return dto.SearchDto{}, nil
	}

	return resp, nil
}

func FetchMovieSearchResults(client *config.TMDBClient, query, page string) (dto.SearchDto, error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := HttpGet[dto.SearchDto](client, endpoints.TmdbEndpoint.Search.Movie, queryParams)
	if err != nil {
		return dto.SearchDto{}, nil
	}

	for i := range resp.Results {
		resp.Results[i].MediaType = "movie"
	}

	return resp, nil
}

func FetchTvSearchResults(client *config.TMDBClient, query, page string) (dto.SearchDto, error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := HttpGet[dto.SearchDto](client, endpoints.TmdbEndpoint.Search.TV, queryParams)
	if err != nil {
		return dto.SearchDto{}, nil
	}

	for i := range resp.Results {
		resp.Results[i].MediaType = "tv"
	}

	return resp, nil
}

func FetchPersonSearchResults(client *config.TMDBClient, query, page string) (dto.SearchDto, error) {
	queryParams := map[string]string{
		"query": query,
		"page":  page,
	}

	resp, err := HttpGet[dto.SearchDto](client, endpoints.TmdbEndpoint.Search.Person, queryParams)
	if err != nil {
		return dto.SearchDto{}, nil
	}

	for i := range resp.Results {
		resp.Results[i].MediaType = "person"
	}

	return resp, nil
}

func FetchSearchWithFilters(client *config.TMDBClient, params types.FilterParams, media string) (dto.SearchDto, error) {
	var release string
	if media == "movie" {
		release = "primary_release_date"
	} else if media == "tv" {
		release = "first_air_date"
	} else {
		return dto.SearchDto{}, fmt.Errorf("invalid media type")
	}

	queryParams := map[string]string{
		"page":                         params.Page,
		"with_watch_providers":         params.Providers,
		"with_genres":                  params.Genres,
		fmt.Sprintf("%s.gte", release): fmt.Sprintf("%s-01-01", params.From),
		fmt.Sprintf("%s.lte", release): fmt.Sprintf("%s-12-31", params.To),
		"sort_by":                      params.Sort,
		"without_genres":               "10763,10764,10767",
		"vote_count.gte":               "300",
		"with_runtime.gte":             params.RuntimeGte,
		"with_runtime.lte":             params.RuntimeLte,
		"watch_region":                 "IT",
	}

	log.Print("Search with filters: ", queryParams)

	resp, err := HttpGet[dto.SearchDto](client, endpoints.TmdbEndpoint.Discover.All(media), queryParams)
	if err != nil {
		return dto.SearchDto{}, nil
	}

	for i := range resp.Results {
		resp.Results[i].MediaType = media
	}

	return resp, nil
}

package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
	"golang.org/x/sync/errgroup"
)

type ContentInfo struct {
	Title        string    `json:"title"`
	ContentID    string    `json:"content_id"`
	ContentType  string    `json:"content_type"`
	BackdropPath string    `json:"backdrop_path"`
	PosterPath   string    `json:"poster_path"`
	ReleaseDate  string    `json:"release_date"`
	Duration     float64   `json:"duration"`
	Genres       []float64 `json:"genres"`
}

func FetchLandingCards(client *config.TMDBClient) (dto.HomeListsDto, error) {
	var g errgroup.Group
	var trendingMovies, trendingTvShows, topRatedMovies, topRatedTvShows []types.ContentCard

	g.Go(func() error {
		results, err := HttpGet[utils.Response[[]types.ContentCard]](client, endpoints.TmdbEndpoint.Trending.Movies, nil)
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
		results, err := HttpGet[utils.Response[[]types.ContentCard]](client, endpoints.TmdbEndpoint.Trending.TV, nil)
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
		results, err := HttpGet[utils.Response[[]types.ContentCard]](client, endpoints.TmdbEndpoint.TopRated.Movies, nil)
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
		results, err := HttpGet[utils.Response[[]types.ContentCard]](client, endpoints.TmdbEndpoint.TopRated.TV, nil)
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
		return dto.HomeListsDto{}, err
	}

	return dto.HomeListsDto{
		TrendingMovies: trendingMovies,
		TrendingTv:     trendingTvShows,
		TopRatedMovies: topRatedMovies,
		TopRatedTv:     topRatedTvShows,
	}, nil
}

func AddUserContent(client *config.SupabaseClient, userId string, contentData ContentInfo) (int, error) {
	contentEndpoint := "/rest/v1/content"
	contentBody := map[string]any{
		"content_id":    contentData.ContentID,
		"content_type":  contentData.ContentType,
		"title":         contentData.Title,
		"backdrop_path": contentData.BackdropPath,
		"poster_path":   contentData.PosterPath,
		"release_date":  contentData.ReleaseDate,
		"duration":      contentData.Duration,
		"genres":        contentData.Genres,
	}
	contentQuery := map[string]string{
		"on_conflict": "content_id,content_type",
		"select":      "id",
	}

	resp, err := client.Post(contentEndpoint, contentQuery, contentBody)
	if resp == nil {
		return 500, fmt.Errorf("failed to send content request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	var results []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil || len(results) == 0 {
		return 500, fmt.Errorf("failed to decode content response or empty result: %w", err)
	}
	contentPrimaryID, ok := results[0]["id"]
	if !ok {
		return 500, fmt.Errorf("content id not found in response")
	}

	watchEndpoint := "/rest/v1/watch"
	watchBody := map[string]any{
		"id":      contentPrimaryID,
		"user_id": userId,
	}
	watchResp, watchErr := client.Post(watchEndpoint, nil, watchBody)
	if watchResp == nil {
		return 500, fmt.Errorf("failed to send watch request: %w", watchErr)
	}
	if watchErr != nil {
		return watchResp.StatusCode, fmt.Errorf("%w", watchErr)
	}
	defer watchResp.Body.Close()

	return watchResp.StatusCode, nil
}

func AddUserContentWatchlist(client *config.SupabaseClient, userId string, contentData ContentInfo) (int, error) {
	contentEndpoint := "/rest/v1/content"

	contentBody := map[string]any{
		"content_id":    contentData.ContentID,
		"content_type":  contentData.ContentType,
		"title":         contentData.Title,
		"backdrop_path": contentData.BackdropPath,
		"poster_path":   contentData.PosterPath,
		"release_date":  contentData.ReleaseDate,
		"duration":      contentData.Duration,
		"genres":        contentData.Genres,
	}
	contentQuery := map[string]string{
		"on_conflict": "content_id,content_type",
		"select":      "id",
	}

	resp, err := client.Post(contentEndpoint, contentQuery, contentBody)
	if resp == nil {
		return 500, fmt.Errorf("failed to send content request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	var results []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil || len(results) == 0 {
		return 500, fmt.Errorf("failed to decode content response or empty result: %w", err)
	}
	contentPrimaryID, ok := results[0]["id"]
	if !ok {
		return 500, fmt.Errorf("content id not found in response")
	}

	watchEndpoint := "/rest/v1/watchlist"
	watchBody := map[string]any{
		"id":      contentPrimaryID,
		"user_id": userId,
	}
	watchResp, watchErr := client.Post(watchEndpoint, nil, watchBody)
	if watchResp == nil {
		return 500, fmt.Errorf("failed to send watch request: %w", watchErr)
	}
	if watchErr != nil {
		return watchResp.StatusCode, fmt.Errorf("%w", watchErr)
	}
	defer watchResp.Body.Close()

	return watchResp.StatusCode, nil
}

func EditRating(client *config.SupabaseClient, userId string, contentId string, contentType string, rating float64) (int, error) {
	internalId, statusCode, err := getInternalContentId(client, contentId, contentType)
	if err != nil {
		log.Println("Error getting internal content ID:", err)
		return statusCode, err
	}

	endpoint := "/rest/v1/watch"
	body := map[string]any{
		"rating": rating,
	}

	queryParams := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"id":      fmt.Sprintf("eq.%.0f", internalId),
	}

	log.Println("Editing rating for internal content ID:", internalId, "for user:", userId, "with rating:", rating)

	resp, err := client.Patch(endpoint, queryParams, body)
	if resp == nil {
		log.Println("Error editing rating: response is nil, returning 500, error:", err)
		return 500, fmt.Errorf("failed to send request: %w", err)
	}
	if err != nil {
		log.Println("Error editing rating:", err)
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func DeleteUserContent(client *config.SupabaseClient, userId string, contentType string, contentId string) (int, error) {
	internalId, statusCode, err := getInternalContentId(client, contentId, contentType)
	if err != nil {
		return statusCode, err
	}

	endpoint := "/rest/v1/watch"
	queryParams := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"id":      fmt.Sprintf("eq.%.0f", internalId),
	}

	resp, err := client.Delete(endpoint, queryParams, nil)
	if resp == nil {
		return 500, fmt.Errorf("failed to send request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func DeleteUserContentWatchlist(client *config.SupabaseClient, userId string, contentType string, contentId string) (int, error) {
	internalId, statusCode, err := getInternalContentId(client, contentId, contentType)
	if err != nil {
		return statusCode, err
	}

	endpoint := "/rest/v1/watchlist"
	queryParams := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"id":      fmt.Sprintf("eq.%.0f", internalId),
	}

	resp, err := client.Delete(endpoint, queryParams, nil)
	if resp == nil {
		return 500, fmt.Errorf("failed to send request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func getInternalContentId(client *config.SupabaseClient, contentId string, contentType string) (any, int, error) {
	endpoint := "/rest/v1/content"
	query := map[string]string{
		"content_id":   fmt.Sprintf("eq.%s", contentId),
		"content_type": fmt.Sprintf("eq.%s", contentType),
		"select":       "id",
	}

	resp, err := client.Get(endpoint, query)
	if resp == nil {
		return nil, 500, fmt.Errorf("failed to get content: %w", err)
	}
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to get content: %w", err)
	}
	defer resp.Body.Close()

	var results []map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, 500, fmt.Errorf("failed to decode content response: %w", err)
	}

	if len(results) == 0 {
		return nil, 404, fmt.Errorf("content not found")
	}

	internalId, ok := results[0]["id"]
	if !ok {
		return nil, 500, fmt.Errorf("content id not found in response")
	}

	return internalId, 200, nil
}

package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/riumat/cinehive-be/config"
)

type UserStats struct {
	TotalWatched   int          `json:"total_watched"`
	TotalWatchlist int          `json:"total_watchlist"`
	TopGenres      []GenreStats `json:"top_genres"`
}

type GenreStats struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage"`
	Count      int     `json:"count"`
}

func FetchUserStats(client *config.SupabaseClient, userId string) (*UserStats, error) {
	stats := &UserStats{}

	totalWatched, err := getTotalWatched(client, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get total watched: %w", err)
	}
	stats.TotalWatched = totalWatched

	totalWatchlist, err := getTotalWatchlist(client, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get total watchlist: %w", err)
	}
	stats.TotalWatchlist = totalWatchlist

	topGenres, err := getTopGenres(client, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get top genres: %w", err)
	}
	stats.TopGenres = topGenres

	return stats, nil
}

func getTotalWatched(client *config.SupabaseClient, userId string) (int, error) {
	endpoint := "/rest/v1/watch"
	query := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"select":  "count",
	}

	resp, err := client.Get(endpoint, query)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return len(result), nil
}

func getTotalWatchlist(client *config.SupabaseClient, userId string) (int, error) {
	endpoint := "/rest/v1/watchlist"
	query := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"select":  "count",
	}

	resp, err := client.Get(endpoint, query)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return len(result), nil
}

func getTopGenres(client *config.SupabaseClient, userId string) ([]GenreStats, error) {
	endpoint := "/rest/v1/watch"
	query := map[string]string{
		"user_id": fmt.Sprintf("eq.%s", userId),
		"select":  "content!inner(content_id,content_type,genres)",
	}

	resp, err := client.Get(endpoint, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var watchedContent []struct {
		Content struct {
			ContentID   string        `json:"content_id"`
			ContentType string        `json:"content_type"`
			Genres      []interface{} `json:"genres"`
		} `json:"content"`
	}

	if err := json.Unmarshal(body, &watchedContent); err != nil {
		return nil, err
	}

	genreCount := make(map[int]struct {
		name  string
		count int
	})

	totalGenreOccurrences := 0

	for _, item := range watchedContent {
		for _, genreInterface := range item.Content.Genres {
			if genreMap, ok := genreInterface.(map[string]interface{}); ok {
				if idFloat, ok := genreMap["id"].(float64); ok {
					if nameStr, ok := genreMap["name"].(string); ok {
						genreID := int(idFloat)

						if existing, exists := genreCount[genreID]; exists {
							genreCount[genreID] = struct {
								name  string
								count int
							}{name: existing.name, count: existing.count + 1}
						} else {
							genreCount[genreID] = struct {
								name  string
								count int
							}{name: nameStr, count: 1}
						}
						totalGenreOccurrences++
					}
				}
			}
		}
	}

	var topGenres []GenreStats
	for id, data := range genreCount {
		percentage := 0.0
		if totalGenreOccurrences > 0 {
			percentage = (float64(data.count) / float64(totalGenreOccurrences)) * 100
		}

		topGenres = append(topGenres, GenreStats{
			ID:         id,
			Name:       data.name,
			Count:      data.count,
			Percentage: percentage,
		})
	}

	for i := 0; i < len(topGenres)-1; i++ {
		for j := i + 1; j < len(topGenres); j++ {
			if topGenres[i].Percentage < topGenres[j].Percentage {
				topGenres[i], topGenres[j] = topGenres[j], topGenres[i]
			}
		}
	}

	return topGenres, nil
}

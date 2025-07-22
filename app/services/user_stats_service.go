package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

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
		"select":  "id",
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
		"select":  "id",
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
			ContentID   float64 `json:"content_id"`
			ContentType string  `json:"content_type"`
			Genres      string  `json:"genres"`
		} `json:"content"`
	}

	if err := json.Unmarshal(body, &watchedContent); err != nil {
		log.Println("Error unmarshalling watched content:", err)
		return nil, err
	}

	log.Printf("Found %d watched content items", len(watchedContent))

	genreCount := make(map[int]int)
	totalGenreOccurrences := 0

	for _, item := range watchedContent {
		var genreIDs []int
		if item.Content.Genres != "" {
			if err := json.Unmarshal([]byte(item.Content.Genres), &genreIDs); err != nil {
				log.Printf("Error parsing genres string for content ID %.0f: %v (genres: %s)", item.Content.ContentID, err, item.Content.Genres)
				continue
			}
		}

		for _, genreID := range genreIDs {
			genreCount[genreID]++
			totalGenreOccurrences++
		}
	}

	log.Printf("Total genre occurrences: %d", totalGenreOccurrences)

	if totalGenreOccurrences == 0 {
		return []GenreStats{}, nil
	}

	genreNames := map[int]string{
		28:    "Action",
		12:    "Adventure",
		16:    "Animation",
		35:    "Comedy",
		80:    "Crime",
		99:    "Documentary",
		18:    "Drama",
		10751: "Family",
		14:    "Fantasy",
		36:    "History",
		27:    "Horror",
		10402: "Music",
		9648:  "Mystery",
		10749: "Romance",
		878:   "Science Fiction",
		10770: "TV Movie",
		53:    "Thriller",
		10752: "War",
		37:    "Western",
		// TV Genres
		10759: "Action & Adventure",
		10762: "Kids",
		10763: "News",
		10764: "Reality",
		10765: "Sci-Fi & Fantasy",
		10766: "Soap",
		10767: "Talk",
		10768: "War & Politics",
	}

	var topGenres []GenreStats
	for genreID, count := range genreCount {
		percentage := (float64(count) / float64(totalGenreOccurrences)) * 100

		genreName, exists := genreNames[genreID]
		if !exists {
			genreName = fmt.Sprintf("Unknown Genre (%d)", genreID)
		}

		topGenres = append(topGenres, GenreStats{
			ID:         genreID,
			Name:       genreName,
			Count:      count,
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

	if len(topGenres) > 5 {
		topGenres = topGenres[:5]
	}

	log.Printf("Returning top %d genres", len(topGenres))

	return topGenres, nil
}

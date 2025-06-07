package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
)

type ReturnType struct {
	Watch     bool    `json:"watch"`
	Rating    float64 `json:"rating"`
	Watchlist bool    `json:"watchlist"`
	ID        float64 `json:"id"`
}

func FetchContentUserData(client *config.SupabaseClient, userId string, contentID string, contentType string) (*ReturnType, error) {
	query := map[string]string{
		"content_id":   fmt.Sprintf("eq.%s", contentID),
		"content_type": fmt.Sprintf("eq.%s", contentType),
		"select":       "id,watch(user_id,rating),watchlist(user_id)",
	}

	resp, err := client.Get(endpoints.Supabase.Tables.Content, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var contentResults []struct {
		ID    any `json:"id"`
		Watch []struct {
			UserID string   `json:"user_id"`
			Rating *float64 `json:"rating"`
		} `json:"watch"`
		Watchlist []struct {
			UserID string `json:"user_id"`
		} `json:"watchlist"`
	}
	if err := json.Unmarshal(body, &contentResults); err != nil {
		return nil, err
	}
	if len(contentResults) == 0 {
		return &ReturnType{Watch: false, Rating: 0, Watchlist: false}, nil
	}

	var watched, watchlisted bool
	var rating float64

	// Cerca userId tra i risultati delle join
	for _, w := range contentResults[0].Watch {
		if w.UserID == userId {
			watched = true
			if w.Rating != nil {
				rating = *w.Rating
			}
			break
		}
	}
	for _, wl := range contentResults[0].Watchlist {
		if wl.UserID == userId {
			watchlisted = true
			break
		}
	}

	return &ReturnType{
		Watch:     watched,
		Rating:    rating,
		Watchlist: watchlisted,
		ID:        contentResults[0].ID.(float64),
	}, nil
}

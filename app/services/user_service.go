package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"golang.org/x/sync/errgroup"
)

type ReturnType struct {
	Watch     bool  `json:"watch"`
	Rating    *int8 `json:"rating"`
	Watchlist bool  `json:"watchlist"`
}

func FetchContentUserData(client *config.SupabaseClient, userId string, contentID float64, contentType string) (*ReturnType, error) {
	queryParams := map[string]string{
		"user_id":      fmt.Sprintf("eq.%s", userId),
		"content_id":   fmt.Sprintf("eq.%s", strconv.FormatFloat(contentID, 'f', -1, 64)),
		"content_type": fmt.Sprintf("eq.%s", contentType),
		"select":       "*",
	}

	var watched bool
	var watchlisted bool
	var rating *int8

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		resp, err := client.Get(endpoints.Supabase.Tables.Content, queryParams)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var results []any
		if err := json.Unmarshal(body, &results); err != nil {
			return err
		}

		if len(results) > 0 {
			watched = true
			log.Println("FetchUserMovieData content results:", results)
			r := results[0].(map[string]any)["rating"]
			if r != nil {
				rating = r.(*int8)
			} else {
				rating = nil
			}
		} else {
			watched = false
		}

		return nil
	})

	g.Go(func() error {
		resp, err := client.Get(endpoints.Supabase.Tables.Watchlist, queryParams)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var results []any
		if err := json.Unmarshal(body, &results); err != nil {
			return err
		}

		log.Println("FetchUserMovieData watchlist results:", results)

		if len(results) > 0 {
			watchlisted = true
		} else {
			watchlisted = false
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	res := ReturnType{
		Watch:     watched,
		Rating:    rating,
		Watchlist: watchlisted,
	}

	log.Println("FetchUserMovieData response:", res)
	return &res, nil
}

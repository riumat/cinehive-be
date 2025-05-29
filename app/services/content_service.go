package services

import (
	"fmt"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/types"
	"golang.org/x/sync/errgroup"
)

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

func AddUserContent(client *config.SupabaseClient, userId string, contentId float64, mediaType string) (int, error) {
	endpoint := "/rest/v1/content"

	body := map[string]any{
		"user_id":      userId,
		"content_id":   contentId,
		"content_type": mediaType,
	}

	resp, err := client.Post(endpoint, nil, body)
	if resp == nil {
		return 500, fmt.Errorf("failed to send request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

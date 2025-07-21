package services

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
	"github.com/riumat/cinehive-be/pkg/utils/helpers"
	"golang.org/x/sync/errgroup"
)

const (
	TV = "tv"
)

func FetchTvDetails(client *config.TMDBClient, id string) (dto.TvDetailsDto, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(TV, id, []string{"external_ids", "aggregate_credits", "recommendations", "videos", "watch%2Fproviders"}), nil)
	if err != nil {
		return dto.TvDetailsDto{}, err
	}

	if prov, ok := results["watch/providers"].(map[string]any); ok {
		if res, ok := prov["results"].(map[string]any); ok {
			if it, ok := res["IT"].(map[string]any); ok {
				results["providers_link"] = it["link"]
			}
		}
	}

	if recs, ok := results["recommendations"].(map[string]any); ok {
		results["recommendations"] = recs["results"]
	}
	if videos, ok := results["videos"].(map[string]any); ok {
		trailers, _ := utils.FormatVideoList(videos["results"].([]any))
		results["videos"] = trailers
	}

	if credits, ok := results["aggregate_credits"].(map[string]any); ok {
		if castArr, ok := credits["cast"].([]any); ok {
			credits["cast"] = helpers.ExtractCastItems(castArr)
		}
		if crewArr, ok := credits["crew"].([]any); ok {
			credits["crew"] = utils.FormatCrewTvList(crewArr)
		}
		results["credits"] = credits
	}

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return dto.TvDetailsDto{}, err
	}

	var resp dto.TvDetailsDto
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return dto.TvDetailsDto{}, err
	}

	return resp, nil
}

func FetchTotalTvDuration(contentId string, seasons int) (float64, error) {

	var (
		totalDuration float64
		mu            sync.Mutex
	)
	g, _ := errgroup.WithContext(context.Background())

	for seasonNum := 1; seasonNum <= seasons; seasonNum++ {
		seasonNum := seasonNum
		type seasonData struct {
			Episodes []struct {
				Runtime float64 `json:"runtime"`
			} `json:"episodes"`
		}

		g.Go(func() error {
			seasonResp, err := HttpGet[seasonData](
				config.NewTMDBClient(),
				endpoints.TmdbEndpoint.DynamicContent.Season(TV, contentId, fmt.Sprintf("%d", seasonNum)),
				nil,
			)
			if err != nil {
				return err
			}

			var seasonDuration float64
			for _, ep := range seasonResp.Episodes {
				seasonDuration += ep.Runtime
			}
			mu.Lock()
			totalDuration += seasonDuration
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}
	return totalDuration, nil
}

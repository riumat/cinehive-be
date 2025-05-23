package services

import (
	"encoding/json"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
)

const (
	MOVIE = "movie"
)

func FetchMovieDetails(client *config.TMDBClient, id string) (dto.MovieDetailsDto, error) {
	results, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.DynamicContent.AllWithAppend(MOVIE, id, []string{"external_ids", "credits", "recommendations", "videos", "watch%2Fproviders"}), nil)
	if err != nil {
		return dto.MovieDetailsDto{}, err
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

	if credits, ok := results["credits"].(map[string]any); ok {
		if crewArr, ok := credits["crew"].([]any); ok {
			credits["crew"] = utils.FormatMovieCrewList(crewArr)
		}
	}

	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return dto.MovieDetailsDto{}, err
	}

	var resp dto.MovieDetailsDto
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return dto.MovieDetailsDto{}, err
	}

	return resp, nil
}

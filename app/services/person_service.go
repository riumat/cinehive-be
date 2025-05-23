package services

import (
	"encoding/json"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
)

func FetchPersonDetails(client *config.TMDBClient, id string) (dto.PersonDto, error) {
	appends := []string{
		"external_ids",
		"combined_credits",
	}

	data, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.Person.AllWithAppend(id, appends), nil)
	if err != nil {
		return dto.PersonDto{}, err
	}

	creditsData, _ := data["combined_credits"].(map[string]any)
	externalData := data["external_ids"]

	var castCredits, crewCredits []map[string]any
	if creditsData != nil {
		if cast, ok := creditsData["cast"].([]any); ok {
			for _, c := range cast {
				if m, ok := c.(map[string]any); ok {
					castCredits = append(castCredits, m)
				}
			}
		}
		if crew, ok := creditsData["crew"].([]any); ok {
			for _, c := range crew {
				if m, ok := c.(map[string]any); ok {
					crewCredits = append(crewCredits, m)
				}
			}
		}
	}

	knownForDepartment, _ := data["known_for_department"].(string)
	var knownForCredits []map[string]any
	if knownForDepartment == "Acting" {
		knownForCredits = utils.FormatCombinedCredits(castCredits)
	} else {
		knownForCredits = utils.FormatCombinedCredits(crewCredits)
	}

	formattedCastCredits := utils.FormatCreditsReleaseDate(castCredits)
	formattedCrewCredits := utils.FormatCreditsReleaseDate(crewCredits)

	result := map[string]any{}
	for k, v := range data {
		result[k] = v
	}
	result["known_for"] = knownForCredits
	result["external_ids"] = externalData
	result["cast_credits"] = formattedCastCredits
	result["crew_credits"] = formattedCrewCredits

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return dto.PersonDto{}, err
	}

	var resp dto.PersonDto
	if err := json.Unmarshal(jsonBytes, &resp); err != nil {
		return dto.PersonDto{}, err
	}

	return resp, nil
}

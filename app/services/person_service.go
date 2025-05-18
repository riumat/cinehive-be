package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/utils"
)

func FetchPersonDetails(client *config.TMDBClient, id string) (any, error) {
	appends := []string{
		"external_ids",
		"images",
		"combined_credits",
	}

	data, err := HttpGet[map[string]any](client, endpoints.TmdbEndpoint.Person.AllWithAppend(id, appends), nil)
	if err != nil {
		return nil, err
	}

	personData := data
	creditsData, _ := data["combined_credits"].(map[string]any)
	imagesData := data["images"]
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

	knownForDepartment, _ := personData["known_for_department"].(string)
	var knownForCredits []map[string]any
	if knownForDepartment == "Acting" {
		knownForCredits = utils.FormatCombinedCredits(castCredits)
	} else {
		knownForCredits = utils.FormatCombinedCredits(crewCredits)
	}

	formattedCastCredits := utils.FormatCreditsReleaseDate(castCredits)
	formattedCrewCredits := utils.FormatCreditsReleaseDate(crewCredits)

	result := map[string]any{}
	for k, v := range personData {
		result[k] = v
	}
	result["combined_credits"] = knownForCredits
	result["images"] = imagesData
	result["external_ids"] = externalData
	result["cast_credits"] = formattedCastCredits
	result["crew_credits"] = formattedCrewCredits

	return result, nil
}

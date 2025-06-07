package services

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils"
)

type PersonInfo struct {
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
	Id          string `json:"id"`
}

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

type ReturnPersonUser struct {
	Follow bool `json:"watch"`
}

func FetchPersonUserData(client *config.SupabaseClient, userId string, personID string) (*ReturnPersonUser, error) {
	query := map[string]string{
		"id":      fmt.Sprintf("eq.%s", personID),
		"user_id": fmt.Sprintf("eq.%s", userId),
		"select":  "id",
	}

	resp, err := client.Get(endpoints.Supabase.Tables.Person, query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var contentResults []struct {
		ID any `json:"id"`
	}
	if err := json.Unmarshal(body, &contentResults); err != nil {
		return nil, err
	}
	if len(contentResults) == 0 {
		return &ReturnPersonUser{Follow: false}, nil
	}

	return &ReturnPersonUser{
		Follow: true,
	}, nil
}

func AddPerson(client *config.SupabaseClient, userId string, personData PersonInfo) (int, error) {
	endpoint := endpoints.Supabase.Tables.Person

	body := map[string]any{
		"user_id":      userId,
		"person_id":    personData.Id,
		"name":         personData.Name,
		"profile_path": personData.ProfilePath,
	}

	resp, err := client.Post(endpoint, nil, body)
	if resp == nil {
		return 500, fmt.Errorf("failed to send request: %w", err)
	}
	if err != nil {
		return resp.StatusCode, fmt.Errorf("%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return resp.StatusCode, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return resp.StatusCode, nil
}

func DeletePerson(client *config.SupabaseClient, userId string, personId string) (int, error) {
	endpoint := endpoints.Supabase.Tables.Person

	queryParams := map[string]string{
		"user_id":   fmt.Sprintf("eq.%s", userId),
		"person_id": fmt.Sprintf("eq.%s", personId),
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

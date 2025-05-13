package utils

import (
	"slices"
	"sort"
	"strings"
)

type CrewItem struct {
	ID          float64 `json:"id"`
	Name        string  `json:"name"`
	ProfilePath string  `json:"profile_path"`
	Job         string  `json:"job"`
}

func FormatCrewList(crew []any) []CrewItem {
	var filteredCrew []any

	// filter
	for _, member := range crew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		job, ok := memberMap["job"].(string)
		if !ok {
			continue
		}

		if slices.Contains(RelevantJobs, job) {
			filteredCrew = append(filteredCrew, member)
		}
	}

	// combine
	crewMap := make(map[string]CrewItem)
	for _, member := range filteredCrew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		name, ok := memberMap["name"].(string)

		if !ok {
			continue
		}
		id, _ := memberMap["id"].(float64)
		profilePath, _ := memberMap["profile_path"].(string)
		job, _ := memberMap["job"].(string)

		if existingMember, exists := crewMap[name]; exists {
			if existingMember.Job != "" {
				existingMember.Job += ", " + job
			} else {
				existingMember.Job = job
			}
			crewMap[name] = existingMember
		} else {
			crewMap[name] = CrewItem{
				ID:          id,
				Name:        name,
				ProfilePath: profilePath,
				Job:         job,
			}
		}
	}

	var formattedCrew []CrewItem
	for _, member := range crewMap {
		formattedCrew = append(formattedCrew, member)
	}

	// sort
	sort.Slice(formattedCrew, func(i, j int) bool {
		jobA := strings.Contains(strings.ToLower(formattedCrew[i].Job), "director")
		jobB := strings.Contains(strings.ToLower(formattedCrew[j].Job), "director")
		if jobA && !jobB {
			return true
		}
		if !jobA && jobB {
			return false
		}

		return false
	})

	return formattedCrew
}

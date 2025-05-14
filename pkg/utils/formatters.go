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

type VideoItem struct {
	ID       float64 `json:"id"`
	Key      string  `json:"key"`
	Name     string  `json:"name"`
	Official bool    `json:"official"`
	Site     string  `json:"site"`
	Type     string  `json:"type"`
	Size     float64 `json:"size"`
}

func FormatVideoList(videos []any) ([]any, []any) {
	var trailers []any
	var others []any

	for _, video := range videos {
		videoMap, ok := video.(map[string]any)
		if !ok {
			continue
		}

		official, ok := videoMap["official"].(bool)
		if !ok || !official {
			continue
		}

		site, ok := videoMap["site"].(string)
		if !ok || (site != "YouTube") {
			continue
		}

		videoType, ok := videoMap["type"].(string)
		if !ok {
			continue
		}

		if videoType == "Trailer" || videoType == "Teaser" {
			trailers = append(trailers, video)
		} else {
			others = append(others, video)
		}
	}

	return trailers, others
}

func FormatMovieCrewList(crew []any) []CrewItem {
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

func FormatCrewTvList(crew []any) []CrewItem {
	var filteredCrew []any

	// filter
	for _, member := range crew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		jobs, ok := memberMap["jobs"].([]any)
		if !ok {
			continue
		}

		for _, jobItem := range jobs {
			jobMap, ok := jobItem.(map[string]any)
			if !ok {
				continue
			}

			job, ok := jobMap["job"].(string)
			if !ok {
				continue
			}

			if slices.Contains(RelevantJobs, job) {
				filteredCrew = append(filteredCrew, member)
				break
			}
		}
	}

	for i, member := range filteredCrew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		jobs, ok := memberMap["jobs"].([]any)
		if !ok {
			continue
		}

		var jobStrings []string
		for _, jobItem := range jobs {
			jobMap, ok := jobItem.(map[string]any)
			if !ok {
				continue
			}

			job, ok := jobMap["job"].(string)
			if !ok {
				continue
			}

			jobStrings = append(jobStrings, job)
		}

		memberMap["job"] = strings.Join(jobStrings, ", ")
		filteredCrew[i] = memberMap
	}

	var formattedCrew []CrewItem
	for _, member := range filteredCrew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		id, _ := memberMap["id"].(float64)
		name, _ := memberMap["name"].(string)
		profilePath, _ := memberMap["profile_path"].(string)
		job, _ := memberMap["job"].(string)

		formattedCrew = append(formattedCrew, CrewItem{
			ID:          id,
			Name:        name,
			ProfilePath: profilePath,
			Job:         job,
		})
	}

	return formattedCrew
}

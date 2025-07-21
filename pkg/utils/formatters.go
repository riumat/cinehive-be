package utils

import (
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/riumat/cinehive-be/pkg/utils/types"
)

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

func FormatMovieCrewList(crew []any) []types.CrewMember {
	var filteredCrew []any

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

	var formattedCrew []types.CrewMember
	for _, member := range filteredCrew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		id, _ := memberMap["id"].(float64)
		name, _ := memberMap["name"].(string)
		profilePath, _ := memberMap["profile_path"].(string)
		job, _ := memberMap["job"].(string)
		department, _ := memberMap["department"].(string)
		popularity, _ := memberMap["popularity"].(float64)

		formattedCrew = append(formattedCrew, types.CrewMember{
			ID:          id,
			Name:        name,
			ProfilePath: profilePath,
			Job:         job,
			Department:  department,
			Popularity:  popularity,
		})
	}

	sort.SliceStable(formattedCrew, func(i, j int) bool {
		return formattedCrew[i].Popularity > formattedCrew[j].Popularity
	})

	sort.Slice(formattedCrew, func(i, j int) bool {
		jobA := strings.ToLower(formattedCrew[i].Job) == "director"
		jobB := strings.ToLower(formattedCrew[j].Job) == "director"
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

func FormatCrewTvList(crew []any) []types.CrewMember {
	var filteredCrew []any

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

	var formattedCrew []types.CrewMember
	for _, member := range filteredCrew {
		memberMap, ok := member.(map[string]any)
		if !ok {
			continue
		}

		id, _ := memberMap["id"].(float64)
		name, _ := memberMap["name"].(string)
		profilePath, _ := memberMap["profile_path"].(string)
		job, _ := memberMap["job"].(string)
		popularity, _ := memberMap["popularity"].(float64)
		department, _ := memberMap["department"].(string)

		formattedCrew = append(formattedCrew, types.CrewMember{
			ID:          id,
			Name:        name,
			ProfilePath: profilePath,
			Job:         job,
			Department:  department,
			Popularity:  popularity,
		})
	}

	sort.Slice(formattedCrew, func(i, j int) bool {
		jobA := strings.ToLower(formattedCrew[i].Job) == "director"
		jobB := strings.ToLower(formattedCrew[j].Job) == "director"
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

func FormatHeaderBackdrop(credits []map[string]any) map[string]any {
	var filteredCredits []map[string]any

	for _, credit := range credits {
		mediaType, _ := credit["media_type"].(string)
		if mediaType != "movie" && mediaType != "tv" {
			continue
		}

		if backdropPath, ok := credit["backdrop_path"].(string); ok && backdropPath != "" {
			voteCount, _ := credit["vote_count"].(float64)

			if character, ok := credit["character"].(string); ok {
				if character == "" || strings.Contains(strings.ToLower(character), "self") {
					continue
				} else {
					log.Println("Non filtrato:", character, credit["name"], credit["title"], credit["popularity"])
				}
			}

			if voteCount > 200 {
				filteredCredits = append(filteredCredits, credit)
			}
		}
	}

	sort.Slice(filteredCredits, func(i, j int) bool {
		popI, _ := filteredCredits[i]["popularity"].(float64)
		popJ, _ := filteredCredits[j]["popularity"].(float64)
		return popI > popJ
	})

	if len(filteredCredits) > 0 {
		return filteredCredits[0]
	}

	return nil
}

func FormatCombinedCredits(cast []map[string]any) []map[string]any {
	excludedGenres := []int{99, 10767, 10764, 10763, 10762, 10768}
	var filteredCredits []map[string]any

	for _, credit := range cast {
		mediaType, _ := credit["media_type"].(string)
		if mediaType != "movie" && mediaType != "tv" {
			continue
		}

		genreIDsAny, ok := credit["genre_ids"].([]any)
		if !ok {
			continue
		}
		var hasExcludedGenre bool
		for _, g := range genreIDsAny {
			genreID, ok := g.(float64)
			if !ok {
				continue
			}
			for _, ex := range excludedGenres {
				if int(genreID) == ex {
					hasExcludedGenre = true
					break
				}
			}
			if hasExcludedGenre {
				break
			}
		}
		if hasExcludedGenre {
			continue
		}

		if mediaType == "tv" {
			episodeCount, _ := credit["episode_count"].(float64)
			if episodeCount < 2 {
				continue
			}
		}

		if character, ok := credit["character"].(string); ok {
			if strings.Contains(strings.ToLower(character), "self") {
				continue
			}
		}

		if mediaType == "movie" {
			order, _ := credit["order"].(float64)
			if order > 10 {
				continue
			}
		}

		voteCount, _ := credit["vote_count"].(float64)
		voteAverage, _ := credit["vote_average"].(float64)
		if voteCount > 100 && voteAverage < 6 {
			continue
		}

		filteredCredits = append(filteredCredits, credit)
	}

	sort.Slice(filteredCredits, func(i, j int) bool {
		vi, _ := filteredCredits[i]["vote_count"].(float64)
		vj, _ := filteredCredits[j]["vote_count"].(float64)
		return vi > vj
	})
	if len(filteredCredits) > 20 {
		filteredCredits = filteredCredits[:20]
	}

	sort.Slice(filteredCredits, func(i, j int) bool {
		vi, _ := filteredCredits[i]["vote_average"].(float64)
		vj, _ := filteredCredits[j]["vote_average"].(float64)
		return vi > vj
	})

	unique := make([]map[string]any, 0, len(filteredCredits))
	seen := make(map[float64]bool)
	for _, credit := range filteredCredits {
		id, _ := credit["id"].(float64)
		if !seen[id] {
			unique = append(unique, credit)
			seen[id] = true
		}
		if len(unique) == 20 {
			break
		}
	}

	return unique
}

func FormatCreditsReleaseDate(list []map[string]any) []map[string]any {
	var result []map[string]any

	for _, credit := range list {
		mediaType, _ := credit["media_type"].(string)
		newCredit := make(map[string]any)
		for k, v := range credit {
			newCredit[k] = v
		}
		if mediaType == "tv" {
			if firstAirDate, ok := credit["first_air_date"].(string); ok {
				newCredit["release_date"] = firstAirDate
			}
		}
		if backdrop, ok := credit["backdrop_path"]; ok {
			newCredit["backdrop_path"] = backdrop
		}
		result = append(result, newCredit)
	}

	var filtered []map[string]any
	for _, credit := range result {
		releaseDate, _ := credit["release_date"].(string)
		if releaseDate != "" {
			filtered = append(filtered, credit)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		dateI, _ := filtered[i]["release_date"].(string)
		dateJ, _ := filtered[j]["release_date"].(string)
		return dateI > dateJ
	})

	return filtered
}

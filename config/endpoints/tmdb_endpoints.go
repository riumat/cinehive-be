package endpoints

import (
	"fmt"
	"strings"
)

const (
	TRENDING        = "/trending"
	MOVIE           = "/movie"
	SEARCH          = "/search"
	PERSON          = "/person"
	DISCOVER        = "/discover"
	GENRE           = "/genre"
	WATCH_PROVIDERS = "/watch/providers"
)

var TmdbEndpoint = struct {
	Trending struct {
		Movies string
		TV     string
	}
	Movie struct {
		Videos func(id int) string
	}
	TopRated struct {
		Movies string
		TV     string
	}
	Search struct {
		Multi  string
		Movie  string
		TV     string
		Person string
	}
	Upcoming struct {
		Movies string
		TV     string
	}
	DynamicContent struct {
		All             func(contentType, id string) string
		AllWithAppend   func(contentType, id string, append []string) string
		Images          func(contentType, id string) string
		Providers       func(contentType, id string) string
		Credits         func(contentType, id, creditType string) string
		Recommendations func(contentType, id string) string
		Videos          func(contentType, id string) string
		Season          func(contentType, id string, seasonNumber string) string
	}
	Person struct {
		All           func(id string) string
		AllWithAppend func(id string, append []string) string
	}
	Discover struct {
		All func(contentType string) string
	}
	Genre struct {
		All func(contentType string) string
	}
	WatchProviders struct {
		All func(contentType string) string
	}
}{
	Trending: struct {
		Movies string
		TV     string
	}{
		Movies: TRENDING + "/movie/week",
		TV:     TRENDING + "/tv/week",
	},
	TopRated: struct {
		Movies string
		TV     string
	}{
		Movies: DISCOVER + "/movie?sort_by=popularity.desc&vote_average.gte=8&vote_count.gte=500",
		TV:     DISCOVER + "/tv?sort_by=popularity.desc&vote_average.gte=8&vote_count.gte=500",
	},
	Movie: struct {
		Videos func(id int) string
	}{
		Videos: func(id int) string {
			return fmt.Sprintf("%s/%d/videos", MOVIE, id)
		},
	},
	Search: struct {
		Multi  string
		Movie  string
		TV     string
		Person string
	}{
		Multi:  SEARCH + "/multi",
		Movie:  SEARCH + "/movie",
		TV:     SEARCH + "/tv",
		Person: SEARCH + "/person",
	},
	Upcoming: struct {
		Movies string
		TV     string
	}{
		Movies: "/movie/upcoming?region=it", //todo collegare queste al service e poi al controller
		TV:     "/tv/on_the_air?timezone=cest",
	},
	DynamicContent: struct {
		All             func(contentType, id string) string
		AllWithAppend   func(contentType, id string, append []string) string
		Images          func(contentType, id string) string
		Providers       func(contentType, id string) string
		Credits         func(contentType, id, creditType string) string
		Recommendations func(contentType, id string) string
		Videos          func(contentType, id string) string
		Season          func(contentType, id string, seasonNumber string) string
	}{
		All: func(contentType, id string) string {
			return fmt.Sprintf("/%s/%s", contentType, id)
		},
		AllWithAppend: func(contentType, id string, append []string) string {
			return fmt.Sprintf("/%s/%s?append_to_response=%s", contentType, id, strings.Join(append, ","))
		},
		Images: func(contentType, id string) string {
			return fmt.Sprintf("/%s/%s/images", contentType, id)
		},
		Providers: func(contentType, id string) string {
			return fmt.Sprintf("/%s/%s/watch/providers", contentType, id)
		},
		Credits: func(contentType, id, creditType string) string {
			return fmt.Sprintf("/%s/%s/%s", contentType, id, creditType)
		},
		Recommendations: func(contentType, id string) string {
			return fmt.Sprintf("/%s/%s/recommendations", contentType, id)
		},
		Videos: func(contentType, id string) string {
			return fmt.Sprintf("/%s/%s/videos", contentType, id)
		},
		Season: func(contentType, id string, seasonNumber string) string {
			return fmt.Sprintf("/%s/%s/season/%s", contentType, id, seasonNumber)
		},
	},
	Person: struct {
		All           func(id string) string
		AllWithAppend func(id string, append []string) string
	}{
		All: func(id string) string {
			return fmt.Sprintf("%s/%s", PERSON, id)
		},
		AllWithAppend: func(id string, append []string) string {
			return fmt.Sprintf("%s/%s?append_to_response=%s", PERSON, id, strings.Join(append, ","))
		},
	},
	Discover: struct {
		All func(contentType string) string
	}{
		All: func(contentType string) string {
			return fmt.Sprintf("%s/%s", DISCOVER, contentType)
		},
	},
	Genre: struct {
		All func(contentType string) string
	}{
		All: func(contentType string) string {
			return fmt.Sprintf("%s/%s/list", GENRE, contentType)
		},
	},
	WatchProviders: struct {
		All func(contentType string) string
	}{
		All: func(contentType string) string {
			return fmt.Sprintf("%s/%s", WATCH_PROVIDERS, contentType)
		},
	},
}

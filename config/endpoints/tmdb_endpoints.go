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
	Search struct {
		Multi string
	}
	DynamicContent struct {
		All             func(contentType, id string) string
		AllWithAppend   func(contentType, id string, append []string) string
		Images          func(contentType, id string) string
		Providers       func(contentType, id string) string
		Credits         func(contentType, id, creditType string) string
		Recommendations func(contentType, id string) string
		Videos          func(contentType, id string) string
	}
	Person struct {
		All      func(id string) string
		Images   func(id string) string
		Credits  func(id string) string
		External func(id string) string
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
	Movie: struct {
		Videos func(id int) string
	}{
		Videos: func(id int) string {
			return fmt.Sprintf("%s/%d/videos", MOVIE, id)
		},
	},
	Search: struct {
		Multi string
	}{
		Multi: SEARCH + "/multi",
	},
	DynamicContent: struct {
		All             func(contentType, id string) string
		AllWithAppend   func(contentType, id string, append []string) string
		Images          func(contentType, id string) string
		Providers       func(contentType, id string) string
		Credits         func(contentType, id, creditType string) string
		Recommendations func(contentType, id string) string
		Videos          func(contentType, id string) string
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
	},
	Person: struct {
		All      func(id string) string
		Images   func(id string) string
		Credits  func(id string) string
		External func(id string) string
	}{
		All: func(id string) string {
			return fmt.Sprintf("%s/%s", PERSON, id)
		},
		Images: func(id string) string {
			return fmt.Sprintf("%s/%s/images", PERSON, id)
		},
		Credits: func(id string) string {
			return fmt.Sprintf("%s/%s/combined_credits", PERSON, id)
		},
		External: func(id string) string {
			return fmt.Sprintf("%s/%s/external_ids", PERSON, id)
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

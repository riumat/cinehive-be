package services

import (
	"fmt"
	"log"

	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
)

func FetchUpcoming(client *config.TMDBClient, media string) (dto.UpcomingResponse, error) {
	var endpoint string
	if media == "movie" {
		endpoint = endpoints.TmdbEndpoint.Upcoming.Movies
	} else if media == "tv" {
		endpoint = endpoints.TmdbEndpoint.Upcoming.TV
	} else {
		return dto.UpcomingResponse{}, fmt.Errorf("invalid media type: %s", media)
	}

	resp, err := HttpGet[dto.UpcomingResponse](client, endpoint, nil)
	if err != nil {
		log.Println("Error fetching upcoming:", err)
		return dto.UpcomingResponse{}, err
	}

	for i := range resp.Results {
		resp.Results[i].MediaType = media
	}

	return resp, nil
}

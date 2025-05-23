package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/config/endpoints"
	"github.com/riumat/cinehive-be/pkg/dto"
	"github.com/riumat/cinehive-be/pkg/utils/types"
)

type CustomResponse struct {
	Genres []types.Genre `json:"genres"`
}

func FetchGenres(client *config.TMDBClient, mediaType string) (dto.GenreDto, error) {
	resp, err := HttpGet[CustomResponse](client, endpoints.TmdbEndpoint.Genre.All(mediaType), nil)

	if err != nil {
		return dto.GenreDto{}, nil
	}

	return resp.Genres, nil
}

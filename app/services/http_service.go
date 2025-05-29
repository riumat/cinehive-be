package services

import (
	"github.com/riumat/cinehive-be/config"
	"github.com/riumat/cinehive-be/pkg/utils"
)

func HttpGet[T any](client *config.TMDBClient, url string, queryParams map[string]string) (T, error) {
	var empty T
	resp, err := client.Get(url, queryParams)
	if err != nil {
		return empty, err
	}

	if err := utils.CheckResponseStatus(resp); err != nil {
		return empty, err
	}

	data, err := utils.DecodeResponseBody[T](resp.Body)
	if err != nil {
		return empty, err
	}

	return data, nil
}

func HttpPost[T any](client *config.TMDBClient, url string, body any, queryParams map[string]string) (T, error) {
	var empty T
	resp, err := client.Post(url, queryParams, body)
	if err != nil {
		return empty, err
	}

	if err := utils.CheckResponseStatus(resp); err != nil {
		return empty, err
	}

	data, err := utils.DecodeResponseBody[T](resp.Body)
	if err != nil {
		return empty, err
	}

	return data, nil
}

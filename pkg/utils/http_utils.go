package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response[T any] struct {
	Results T `json:"results"`
}

type PaginatedResponse[T any] struct {
	Results      T   `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Page         int `json:"page"`
}

func DecodeResponseBody[T any](body io.ReadCloser) (T, error) {
	defer body.Close()

	var response T
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response body: %w", err)
	}

	return response, nil
}

func CheckResponseStatus(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code %d", resp.StatusCode)
	}
	return nil
}

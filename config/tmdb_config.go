package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type TMDBClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewTMDBClient() *TMDBClient {
	return &TMDBClient{
		BaseURL: GetEnv("TMDB_BASE_URL"),
		APIKey:  GetEnv("TMDB_API_KEY"),
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (client *TMDBClient) Get(endpoint string, queryParams map[string]string) (*http.Response, error) {
	return client.request(http.MethodGet, endpoint, queryParams, nil)
}

func (client *TMDBClient) Post(endpoint string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	return client.request(http.MethodPost, endpoint, queryParams, body)
}

func (client *TMDBClient) Put(endpoint string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	return client.request(http.MethodPut, endpoint, queryParams, body)
}

func (client *TMDBClient) Delete(endpoint string, queryParams map[string]string) (*http.Response, error) {
	return client.request(http.MethodDelete, endpoint, queryParams, nil)
}

func (client *TMDBClient) request(method, endpoint string, queryParams map[string]string, body interface{}) (*http.Response, error) {
	fullURL, err := url.Parse(fmt.Sprintf("%s%s", client.BaseURL, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	query := fullURL.Query()
	query.Set("api_key", client.APIKey)
	for key, value := range queryParams {
		query.Set(key, value)
	}
	fullURL.RawQuery = query.Encode()

	var requestBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize body: %w", err)
		}
		requestBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, fullURL.String(), requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	return resp, nil
}

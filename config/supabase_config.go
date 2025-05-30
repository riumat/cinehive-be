package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type SupabaseClient struct {
	BaseURL    string
	ApiKey     string
	HTTPClient *http.Client
}

func NewSupabaseClient(token string) *SupabaseClient {
	return &SupabaseClient{
		BaseURL: os.Getenv("SUPABASE_URL"),
		ApiKey:  token,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (client *SupabaseClient) Get(endpoint string, queryParams map[string]string) (*http.Response, error) {
	return client.request(http.MethodGet, endpoint, queryParams, nil)
}

func (client *SupabaseClient) Post(endpoint string, queryParams map[string]string, body any) (*http.Response, error) {
	return client.request(http.MethodPost, endpoint, queryParams, body)
}

func (client *SupabaseClient) Patch(endpoint string, queryParams map[string]string, body any) (*http.Response, error) {
	return client.request(http.MethodPatch, endpoint, queryParams, body)
}
func (client *SupabaseClient) Delete(endpoint string, queryParams map[string]string, body any) (*http.Response, error) {
	return client.request(http.MethodDelete, endpoint, queryParams, nil)
}

func (client *SupabaseClient) request(method, endpoint string, queryParams map[string]string, body any) (*http.Response, error) {
	fullURL, err := url.Parse(fmt.Sprintf("%s%s", client.BaseURL, endpoint))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if len(queryParams) > 0 {
		q := fullURL.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		fullURL.RawQuery = q.Encode()
	}

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

	headers := map[string]string{
		"apikey": os.Getenv("SUPABASE_ANON_KEY"),
		"Prefer": "return=minimal",
	}
	if client.ApiKey != "" {
		headers["Authorization"] = "Bearer " + client.ApiKey
	} else {
		return nil, fmt.Errorf("user bearer token is not set")
	}
	if body != nil {
		headers["Content-Type"] = "application/json"
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		var supabaseErr struct {
			Message string `json:"message"`
		}
		msg := string(bodyBytes)
		if err := json.Unmarshal(bodyBytes, &supabaseErr); err == nil && supabaseErr.Message != "" {
			msg = supabaseErr.Message
		}
		return resp, fmt.Errorf("%s", msg)
	}

	return resp, nil
}

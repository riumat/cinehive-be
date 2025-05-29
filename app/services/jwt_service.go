package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/riumat/cinehive-be/config/endpoints"
)

type supabaseRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Error        *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func GetNewAuthToken(refreshToken string) (any, error) {
	url := os.Getenv("SUPABASE_URL") + endpoints.Supabase.Auth.Refresh

	bodyData := map[string]string{
		"refresh_token": refreshToken,
	}
	bodyBytes, err := json.Marshal(bodyData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "supabase error: ", err
	}
	defer resp.Body.Close()

	var results supabaseRefreshResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return "decode error: ", err
	}
	return results, nil
}

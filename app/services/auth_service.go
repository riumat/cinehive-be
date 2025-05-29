package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/riumat/cinehive-be/config/endpoints"
)

var SUPABASE_URL = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY = os.Getenv("SUPABASE_SERVICE_ROLE_KEY")

type supabaseAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"user"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func SignUpWithSupabase(email, password string) (any, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", SUPABASE_URL, endpoints.Supabase.Auth.SignUp), bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("apikey", SUPABASE_KEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var supaResp supabaseAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&supaResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if supaResp.Error != nil {
		return "", errors.New(supaResp.Error.Message)
	}
	if supaResp.User.ID == "" {
		return "", errors.New("no user id returned from Supabase")
	}

	return supaResp, nil
}

func SignInWithSupabase(email, password string) (any, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", SUPABASE_URL, endpoints.Supabase.Auth.SignIn), bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("apikey", SUPABASE_KEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var supaResp supabaseAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&supaResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return supaResp, nil
}

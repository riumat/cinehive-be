package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/riumat/cinehive-be/config/endpoints"
)

type supabaseAuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	User         *struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	} `json:"user,omitempty"`

	Code      int    `json:"code,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Msg       string `json:"msg,omitempty"`
}

func AddUserToProfile(userId string, username string) error {
	url := os.Getenv("SUPABASE_URL") + endpoints.Supabase.Tables.Profiles
	anonKey := os.Getenv("SUPABASE_ANON_KEY")

	payload := map[string]string{
		"user_id":  userId,
		"username": username,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("apikey", anonKey)
	req.Header.Set("Authorization", "Bearer "+anonKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var respBody any
	json.NewDecoder(resp.Body).Decode(&respBody)

	log.Println("AddUserToProfile response:", respBody)

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("failed to add user to profile: status code %d", resp.StatusCode)
	}

	return nil
}

func SignUpWithSupabase(email string, password string) (supabaseAuthResponse, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := os.Getenv("SUPABASE_URL") + endpoints.Supabase.Auth.SignUp

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var supaResp supabaseAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&supaResp); err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to decode response: %w", err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return supabaseAuthResponse{Code: resp.StatusCode, Msg: supaResp.ErrorCode}, fmt.Errorf("supabase sign up failed: %s", supaResp.Msg)
	}
	if supaResp.User.ID == "" {
		return supabaseAuthResponse{Code: 500}, errors.New("no user id returned from Supabase")
	}

	return supaResp, nil
}

func SignInWithSupabase(email, password string) (supabaseAuthResponse, error) {
	payload := map[string]string{
		"email":    email,
		"password": password,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := os.Getenv("SUPABASE_URL") + endpoints.Supabase.Auth.SignIn

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var supaResp supabaseAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&supaResp); err != nil {
		return supabaseAuthResponse{Code: 500}, fmt.Errorf("failed to decode response: %w", err)
	}
	return supaResp, nil
}

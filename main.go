package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Auth struct {
	Auth  string `json:"auth"`
	Email string `json:"email"`
}

type AuthsResponse struct {
	Auths map[string]Auth `json:"auths"`
}

func enVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func getBearerToken() (string, error) {
	offlineToken := enVar("OFFLINE_ACCESS_TOKEN")
	if offlineToken == "" {
		return "", fmt.Errorf("OFFLINE_ACCESS_TOKEN is not set in .env")
	}
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", "cloud-services")
	data.Set("refresh_token", offlineToken)
	resp, err := http.PostForm("https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token", data)
	if err != nil {
		return "", fmt.Errorf("failed to get Bearer token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token, received non-200 response: %d, body: %s", resp.StatusCode, string(body))
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %v", err)
	}
	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in the response")
	}
	return accessToken, nil
}

func makeAPIRequest(bearerToken string) error {
	url := "https://api.openshift.com/api/accounts_mgmt/v1/access_token"
	jsonPayload := []byte(`{}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d, body: %s", resp.StatusCode, string(body))
	}
	var result AuthsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}
	jsonOutput, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}
	fmt.Println(string(jsonOutput))
	return nil
}

func main() {
	bearerToken, err := getBearerToken()
	if err != nil {
		fmt.Println("Error getting Bearer token:", err)
		return
	}
	err = makeAPIRequest(bearerToken)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
}

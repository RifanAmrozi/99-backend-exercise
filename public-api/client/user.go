package client

import (
	"99-backend-exercise/models"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func CreateUser(name string) (models.User, error) {
	apiURL := "http://localhost:8082/users"
	data := url.Values{}
	data.Set("name", name)

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return models.User{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.User{}, err
	}
	defer resp.Body.Close()

	var response struct {
		Result bool        `json:"result"`
		User   models.User `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.User{}, err
	}

	return response.User, nil
}

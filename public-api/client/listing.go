package client

import (
	"99-backend-exercise/models"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func CreateListing(userID int, listingType string, price int) (models.Listing, error) {
	apiURL := "http://localhost:8081/listings"
	data := url.Values{}
	data.Set("user_id", strconv.Itoa(userID))
	data.Set("listing_type", listingType)
	data.Set("price", strconv.Itoa(price))

	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return models.Listing{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Listing{}, err
	}
	defer resp.Body.Close()

	var response struct {
		Result  bool           `json:"result"`
		Listing models.Listing `json:"listing"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return models.Listing{}, err
	}

	return response.Listing, nil
}

func GetListings(pageNum, pageSize int, userID *int) ([]models.Listing, error) {
	url := "http://localhost:8081/listings?page_num=" + strconv.Itoa(pageNum) + "&page_size=" + strconv.Itoa(pageSize)
	if userID != nil {
		url += "&user_id=" + strconv.Itoa(*userID)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Result   bool             `json:"result"`
		Listings []models.Listing `json:"listings"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Listings, nil
}

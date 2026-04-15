package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationAreaResponse{}, fmt.Errorf("error unmarshaling data from cache: %w", err)
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error creating request to location area: %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error making request to location area: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error transforming body to bytes: %w", err)
	}

	locationResp := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error unmarshaling data: %w", err)
	}

	c.cache.Add(url, data)
	return locationResp, nil
}

func (c *Client) SearchLocation(locName string) (LocationResponse, error) {
	url := baseURL + "/location-area/" + locName + "/"

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationResponse{}, fmt.Errorf("error unmarshaling data from cache: %w", err)
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error creating request to url: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error making request to url: %w", err)
	}
	locationResp := LocationResponse{}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error transforming body to bytes: %w", err)
	}

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error unmarshaling data: %w", err)
	}

	c.cache.Add(url, data)

	return locationResp, nil
}

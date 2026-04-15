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

	return locationResp, nil
}

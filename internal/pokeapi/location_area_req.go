package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(data, &locationAreasResponse)

		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	c.cache.Add(fullURL, data)
	return locationAreasResponse, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaResponse := LocationArea{}
		err := json.Unmarshal(data, &locationAreaResponse)
		if err != nil {
			return LocationArea{}, err
		}

		return locationAreaResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResponse := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, data)
	return locationAreaResponse, nil
}

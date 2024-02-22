package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		locationAreaRes := LocationAreaRes{}

		jsonErr := json.Unmarshal(data, &locationAreaRes)
		if jsonErr != nil {
			return LocationAreaRes{}, jsonErr
		}

		return locationAreaRes, nil
	}

	fmt.Println("cache miss!")


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaRes{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreaRes{}, fmt.Errorf("bad status code, %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaRes{}, err
	}

	locationAreaRes := LocationAreaRes{}

	jsonErr := json.Unmarshal(data, &locationAreaRes)
	if jsonErr != nil {
		return LocationAreaRes{}, jsonErr
	}

	c.cache.Add(fullURL, data)

	return locationAreaRes, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		locationArea := LocationArea{}

		jsonErr := json.Unmarshal(data, &locationArea)
		if jsonErr != nil {
			return LocationArea{}, jsonErr
		}

		return locationArea, nil
	}

	fmt.Println("cache miss!")


	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code, %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}

	jsonErr := json.Unmarshal(data, &locationArea)
	if jsonErr != nil {
		return LocationArea{}, jsonErr
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}
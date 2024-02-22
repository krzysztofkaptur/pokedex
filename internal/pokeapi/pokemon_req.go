package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		pokemon := Pokemon{}

		jsonErr := json.Unmarshal(data, &pokemon)
		if jsonErr != nil {
			return Pokemon{}, jsonErr
		}

		return pokemon, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code, %v", res.StatusCode)
	}

	data, ioErr := io.ReadAll(res.Body)
	if ioErr != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}

	jsonErr := json.Unmarshal([]byte(data), &pokemon)
	if jsonErr != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(fullURL, data)

	return pokemon, nil
}
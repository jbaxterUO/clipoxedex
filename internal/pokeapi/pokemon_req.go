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

	data, ok := c.cache.Get(fullURL)
	if ok {
		pokemonResponse := Pokemon{}
		err := json.Unmarshal(data, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResponse := Pokemon{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, data)
	return pokemonResponse, nil
}

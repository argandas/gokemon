package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
)

// A Pokedex returns the names and URL for all pokemon
type Pokedex struct {
	Name    string     `json:"name"`
	URL     string     `json:"resource_uri"`
	Pokemon []*Pokemon `json:"pokemon"`
}

// Return a Pokedex with Pokemon Basic Information (Name and API's URL)
func GetPokedex() (*Pokedex, error) {
	resp, err := http.Get(API_URL + "/api/v1/pokedex/1/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	pokedex := new(Pokedex)
	err = json.NewDecoder(resp.Body).Decode(pokedex)
	if err != nil {
		return nil, err
	}
	// Fix for missing "/" on pokemon URIs
	for index, pokemon := range pokedex.Pokemon {
		if pokemon.URL[:1] != "/" {
			pokedex.Pokemon[index].URL = "/" + pokemon.URL
		}
	}
	return pokedex, nil
}

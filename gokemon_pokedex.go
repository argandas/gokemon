package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Pokedex struct {
	Name string
	Pokemons []Pokemon `json:"pokemon"`
}

func NewPokedex() (*Pokedex, error) {
	resp, err := http.Get("http://pokeapi.co/api/v1/pokedex/1/")
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
	return pokedex, nil
}
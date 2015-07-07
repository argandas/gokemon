package gokemon

import (
	"encoding/json"
	"fmt"
	"strings"
	"errors"
	"net/http"
)

// Post describes a Wordpress post.
type Pokedex struct {
	Name string
	Pokemons map[string]Pokemon
	// Aux array for JSON data
	data []Pokemon `json:"pokemon"`
}

// Get fetches the most recent posts from a Wordpress blog.
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
	// Fill pokemons map
	pokedex.Pokemons = make(map[string]Pokemon)
	for _, pokemon := range pokedex.data {
		pokedex.Pokemons[pokemon.Name] = pokemon
	}
	return pokedex, nil
}

type Pokemon struct {
	Name string
	URL string `json:"resource_uri"`
}

func (poke Pokemon) String() string {
	return fmt.Sprintf("\r\n%s\r\n", strings.Title(poke.Name))
}
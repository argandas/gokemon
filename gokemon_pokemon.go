package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

type Pokemon struct {
	Name string 			`json:"name"`
	ID int 					`json:"national_id"`
	URL string 				`json:"resource_uri"`
	Abilities []*Ability 	`json:"abilities"`
	Types []*Type 			`json:"types"`
	Moves []*Move 			`json:"moves"`
}

// Get detailed pokemon data 
func (pokemon *Pokemon) Get() *Pokemon {
	temp, err := getPokemon(API_URL + pokemon.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print pokemon data in string format
func (poke Pokemon) String() string {
	str := ""
	str += fmt.Sprintf("Pokemon: %s\r\n", strings.Title(poke.Name))
	if poke.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", poke.ID)
	}
	str += fmt.Sprintf("  URL: %s\r\n", poke.URL)
	str += ability2string(poke.Abilities, "Abilities")
	str += type2string(poke.Types, "Types")
	str += move2int(poke.Moves, "Moves")
	return str
}



// Return a Pokemon with Basic Information using pokemon name
func GetPokemon(name string) (*Pokemon, error) {
	name = strings.ToLower(name)
	url := fmt.Sprintf(API_URL + "/api/v1/pokemon/%s/", name)
	return getPokemon(url)
}

// Return a Pokemon with Basic Information using pokemon ID
func GetPokemonByID(id int) (*Pokemon, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/pokemon/%d/", id)
	return getPokemon(url)
}

// Pokemon API request
func getPokemon(url string) (*Pokemon, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	pokemon := new(Pokemon)	
	err = json.NewDecoder(resp.Body).Decode(pokemon) 
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
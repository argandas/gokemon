package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

// An Ability resource represent a single Pokémon ability
type Ability struct {
	ID int 					`json:"id"`
	Name string 			`json:"name"`
	Description string 		`json:"description"`
	URL string 				`json:"resource_uri"`
}

// Get detailed ability data 
func (ability *Ability) Get() *Ability {
	temp, err := getAbility(API_URL + ability.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print ability data in string format
func (ability *Ability) String() string {
	str := ""
	str += fmt.Sprintf("Ability: %s\r\n", strings.Title(ability.Name))
	if ability.Description != "" {
		str += fmt.Sprintf("  Description: %s\r\n", ability.Description)
	}
	str += fmt.Sprintf("  URL: %s\r\n", ability.URL)
	if ability.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", ability.ID)
	}
	return str
}

// Return a Ability with Basic Information using ability ID
func GetAbility(id int) (*Ability, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/ability/%d/", id)
	return getAbility(url)
}

// Ability API request
func getAbility(url string) (*Ability, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	ability := new(Ability)	
	err = json.NewDecoder(resp.Body).Decode(ability) 
	if err != nil {
		return nil, err
	}
	return ability, nil
}

func ability2string(data []*Ability, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}
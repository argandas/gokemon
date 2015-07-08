package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

// A Description resource represent a single Pokémon pokédex description
type Description struct {
	ID int 					`json:"id"`
	Name string 			`json:"name"`
	Description string 		`json:"description"`
	URL string 				`json:"resource_uri"`
	Pokemon *Pokemon 		`json:"pokemon"`
	Games []*Game 			`json:"games"`
}

// Get detailed description data 
func (desc *Description) Get() *Description {
	temp, err := getDescription(API_URL + desc.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print description data in string format
func (desc *Description) String() string {
	str := ""
	str += fmt.Sprintf("Description: %s\r\n", strings.Title(desc.Name))
	str += fmt.Sprintf("  URL: %s\r\n", desc.URL)
	if desc.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", desc.ID)
	}
	if desc.Description != "" {
		str += fmt.Sprintf("  Description: %s\r\n", desc.Description)
	}
	if desc.Pokemon != nil {
		str += fmt.Sprintf("  Pokemon: %s\r\n", desc.Pokemon.Name)
	}
	return str
}

// Return a Description with Basic Information using description ID
func GetDescription(id int) (*Description, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/description/%d/", id)
	return getDescription(url)
}

// Description API request
func getDescription(url string) (*Description, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	desc := new(Description)	
	err = json.NewDecoder(resp.Body).Decode(desc) 
	if err != nil {
		return nil, err
	}
	return desc, nil
}

func description2string(data []*Description, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}
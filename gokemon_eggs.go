package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

// An Egg group resource represent a single Pokémon egg group
type Egg struct {
	ID int 					`json:"id"`
	Name string 			`json:"name"`
	URL string 				`json:"resource_uri"`
	Pokemon []*Pokemon		`json:"pokemon"`
}

// Get detailed egg data 
func (egg *Egg) Get() *Egg {
	temp, err := getEgg(API_URL + egg.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print egg data in string format
func (egg *Egg) String() string {
	str := ""
	str += fmt.Sprintf("Egg: %s\r\n", strings.Title(egg.Name))
	str += fmt.Sprintf("  URL: %s\r\n", egg.URL)
	if egg.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", egg.ID)
	}
	return str
}

// Return a Egg with Basic Information using egg ID
func GetEgg(id int) (*Egg, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/egg/%d/", id)
	return getEgg(url)
}

// Egg API request
func getEgg(url string) (*Egg, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	egg := new(Egg)	
	err = json.NewDecoder(resp.Body).Decode(egg) 
	if err != nil {
		return nil, err
	}
	return egg, nil
}

func egg2string(data []*Egg, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}
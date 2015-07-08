package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

// A Sprite resource represent a single Pokémon Sprite
type Sprite struct {
	ID int 				`json:"id"`
	Name string 		`json:"name"`
	Image string 		`json:"image"`
	URL string 			`json:"resource_uri"`
	Pokemon *Pokemon 	`json:"pokemon"`
}

// Get detailed sprite data 
func (sprite *Sprite) Get() *Sprite {
	temp, err := getSprite(API_URL + sprite.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print sprite data in string format
func (sprite *Sprite) String() string {
	str := ""
	str += fmt.Sprintf("Sprite: %s\r\n", strings.Title(sprite.Name))
	str += fmt.Sprintf("  URL: %s\r\n", sprite.URL)
	if sprite.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", sprite.ID)
	}
	if sprite.Pokemon != nil {
		str += fmt.Sprintf("  Pokemon: %s\r\n", sprite.Pokemon.Name)
	}
	if sprite.Image != "" {
		str += fmt.Sprintf("  Image: %s\r\n", sprite.Image)
	}
	return str
}

// Return a Sprite with Basic Information using sprite ID
func GetSprite(id int) (*Sprite, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/sprite/%d/", id)
	return getSprite(url)
}

// Sprite API request
func getSprite(url string) (*Sprite, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	sprite := new(Sprite)	
	err = json.NewDecoder(resp.Body).Decode(sprite) 
	if err != nil {
		return nil, err
	}
	return sprite, nil
}

func sprite2string(data []*Sprite, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}
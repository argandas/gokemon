package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

// A Game resource represent a single Pokémon game
type Game struct {
	ID int 				`json:"id"`
	Name string 		`json:"name"`
	URL string 			`json:"resource_uri"`
	Release int 		`json:"release_year"`
	Generation int 		`json:"generation"`
}

// Get detailed game data 
func (game *Game) Get() *Game {
	temp, err := getGame(API_URL + game.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print game data in string format
func (game *Game) String() string {
	str := ""
	str += fmt.Sprintf("Game: %s\r\n", strings.Title(game.Name))
	str += fmt.Sprintf("  URL: %s\r\n", game.URL)
	if game.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", game.ID)
	}
	if game.Release != 0 {
		str += fmt.Sprintf("  Release year: %d\r\n", game.Release)
	}
	if game.Generation != 0 {
		str += fmt.Sprintf("  Generation: %d\r\n", game.Generation)
	}
	return str
}

// Return a Game with Basic Information using game ID
func GetGame(id int) (*Game, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/game/%d/", id)
	return getGame(url)
}

// Game API request
func getGame(url string) (*Game, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	game := new(Game)	
	err = json.NewDecoder(resp.Body).Decode(game) 
	if err != nil {
		return nil, err
	}
	return game, nil
}

func game2string(data []*Game, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}
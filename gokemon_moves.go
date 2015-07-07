package gokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"fmt"
	"strings"
)

type Move struct {
	ID int 					`json:"id"`
	Name string 			`json:"name"`
	Description string 		`json:"description"`
	URL string 				`json:"resource_uri"`
	Power int 				`json:"power"`
	PP int 					`json:"pp"`
	Accuracy int 			`json:"accuracy"`
}

// Get detailed move data 
func (move *Move) Get() *Move {
	temp, err := getMove(API_URL + move.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print move data in string format
func (move *Move) String() string {
	str := ""
	str += fmt.Sprintf("Move: %s\r\n", strings.Title(move.Name))
	str += fmt.Sprintf("  URL: %s\r\n", move.URL)
	if move.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", move.ID)
	}
	if move.Power != 0 {
		str += fmt.Sprintf("  Power: %d\r\n", move.Power)
	}
	if move.Accuracy != 0 {
		str += fmt.Sprintf("  Accuracy: %d\r\n", move.Accuracy)
	}
	if move.PP != 0 {
		str += fmt.Sprintf("  PP: %d\r\n", move.PP)
	}
	return str
}

// Return a Move with Basic Information using move ID
func GetMove(id int) (*Move, error) {
	url := fmt.Sprintf(API_URL + "/api/v1/move/%d/", id)
	return getMove(url)
}

// Move API request
func getMove(url string) (*Move, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status) 
	}
	move := new(Move)	
	err = json.NewDecoder(resp.Body).Decode(move) 
	if err != nil {
		return nil, err
	}
	return move, nil
}

func move2string(data []*Move, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	} 
	return ""
}

func move2int(data []*Move, title string) string {
	if len(data) > 0 {
		return fmt.Sprintf("  %s: [%d]\r\n", title, len(data))
	} 
	return ""
}
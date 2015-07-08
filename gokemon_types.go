package gokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// A Type resource represent a single Pokémon type
type Type struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	URL            string  `json:"resource_uri"`
	NoEffect       []*Type `json:"no_effect"`
	Ineffective    []*Type `json:"ineffective"`
	Resistance     []*Type `json:"resistance"`
	SuperEffective []*Type `json:"super_effective"`
	Weakness       []*Type `json:"weakness"`
}

// Get detailed type data
func (t *Type) Get() *Type {
	temp, err := getType(API_URL + t.URL)
	if err != nil {
		return nil
	}
	return temp
}

// Print type data in string format
func (t *Type) String() string {
	str := ""
	str += fmt.Sprintf("Type: %s\r\n", strings.Title(t.Name))
	str += fmt.Sprintf("  URL: %s\r\n", t.URL)
	if t.ID != 0 {
		str += fmt.Sprintf("  ID: %d\r\n", t.ID)
	}
	str += type2string(t.NoEffect, "No-Effect")
	str += type2string(t.Ineffective, "Ineffective")
	str += type2string(t.SuperEffective, "Super-Effective")
	str += type2string(t.Resistance, "Resistance")
	str += type2string(t.Weakness, "Weakness")
	return str
}

// Return a type with Basic Information using type ID
func GetType(id int) (*Type, error) {
	url := fmt.Sprintf(API_URL+"/api/v1/type/%d/", id)
	return getType(url)
}

// Type API request
func getType(url string) (*Type, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	t := new(Type)
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func type2string(data []*Type, title string) string {
	if len(data) > 0 {
		items := []string{}
		for _, item := range data {
			items = append(items, item.Name)
		}
		return fmt.Sprintf("  %s: %s\r\n", title, strings.Join(items, ", "))
	}
	return ""
}

package gokemon

import (
	"fmt"
	"strings"
)

type Pokemon struct {
	Name string
	URL string `json:"resource_uri"`
}

func (poke Pokemon) String() string {
	name := strings.Title(poke.Name)
	for len(name) <= 15 {
		name += " "
	}
	return fmt.Sprintf("%s\t%s", name, poke.URL)
}
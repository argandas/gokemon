package gokemon

import (
	"fmt"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{
		{"Charizard", 6},
		{"Mewtwo", 150},
		{"Gardevoir", 282},
	}
	for _, c := range cases {
		pokemon, err := GetPokemon(c.name)
		if err != nil {
			t.Errorf("GetPokemon(%s) >> Error: %s", c.name, err)
		}
		if pokemon.ID != c.id {
			t.Errorf("GetPokemon(%s) == %d, want %d", c.name, pokemon.ID, c.id)
		} else {
			fmt.Println(pokemon)
		}
	}
}

func TestGetPokemonByID(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{
		{"Tepig", 498},
		{"Duskull", 355},
		{"Caterpie", 10},
	}
	for _, c := range cases {
		pokemon, err := GetPokemonByID(c.id)
		if err != nil {
			t.Errorf("GetPokemonByID(%d) >> Error: %s", c.id, err)
		}
		if pokemon.Name != c.name {
			t.Errorf("GetPokemonByID(%d) == %s, want %s", c.id, pokemon.Name, c.name)
		} else {
			fmt.Println(pokemon)
		}
	}
}

func TestGetAbility(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{
		{"Pride", 200},
		{"Strong-jaw", 185},
		{"Stench", 1},
	}
	for _, c := range cases {
		ability, err := GetAbility(c.id)
		if err != nil {
			t.Errorf("GetAbility(%d) >> Error: %s", c.id, err)
		}
		if ability.Name != c.name {
			t.Errorf("GetAbility(%d) == %s, want %s", c.id, ability.Name, c.name)
		} else {
			fmt.Println(ability)
		}
	}
}

func TestGetType(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{
		{"Grass", 12},
		{"Normal", 1},
		{"Ghost", 8},
	}
	for _, c := range cases {
		pokeType, err := GetType(c.id)
		if err != nil {
			t.Errorf("GetType(%d) >> Error: %s", c.id, err)
		}
		if pokeType.Name != c.name {
			t.Errorf("GetType(%d) == %s, want %s", c.id, pokeType.Name, c.name)
		} else {
			fmt.Println(pokeType)
		}
	}
}

func TestGetMove(t *testing.T) {
	cases := []struct {
		name string
		id   int
	}{
		{"Pound", 1},
		{"Teleport", 100},
		{"Spit-up", 255},
	}
	for _, c := range cases {
		move, err := GetMove(c.id)
		if err != nil {
			t.Errorf("GetMove(%d) >> Error: %s", c.id, err)
		}
		if move.Name != c.name {
			t.Errorf("GetMove(%d) == %s, want %s", c.id, move.Name, c.name)
		} else {
			fmt.Println(move)
		}
	}
}

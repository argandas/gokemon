package main

import (
	"fmt"
	"github.com/argandas/gokemon"
)

func main() {
	// Get Pokedex
	pokedex, err := gokemon.GetPokedex()
	if err != nil {
		panic(err)
	}
	// Print first 3 pokemons
	for _, pokemon := range pokedex.Pokemon[:3] {
		fmt.Println(pokemon.Get())
	}

	// Get info about Blastoise
	blastoise, err := gokemon.GetPokemonByID(9)
	if err != nil {
		panic(err)
	}
	fmt.Println(blastoise)

	// Get info about Charizard
	charizard, err := gokemon.GetPokemon("charizard")
	if err != nil {
		panic(err)
	}
	fmt.Println(charizard)

	// Get Ability
	swarm, err := gokemon.GetAbility(68)
	if err != nil {
		panic(err)
	}
	fmt.Println(swarm)

	// Get Type
	flying, err := gokemon.GetType(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(flying)

	// Get Move
	dragonTail, err := gokemon.GetMove(525)
	if err != nil {
		panic(err)
	}
	fmt.Println(dragonTail)

	// Get Egg
	monster, err := gokemon.GetEgg(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(monster)

	// Get Description
	desc, err := gokemon.GetDescription(200)
	if err != nil {
		panic(err)
	}
	fmt.Println(desc)

	// Get Game
	yellow, err := gokemon.GetGame(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(yellow)

	// Get Sprite
	sprite, err := gokemon.GetSprite(77)
	if err != nil {
		panic(err)
	}
	fmt.Println(sprite)

}
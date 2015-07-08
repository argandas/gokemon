# gokemon
A Golang wrapper for [PokeAPI](http://pokeapi.co/)

## Installation

```bash
	go get github.com/argandas/gokemon
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/argandas/gokemon"
)

func main() {
	// Get info about Charizard
	charizard, err := gokemon.GetPokemon("charizard")
	if err != nil {
		panic(err)
	}
	fmt.Println(charizard)
	// Get Charizard types
	for _, type := range charizard.Types {
		fmt.Println(type)
	}
}
```

### Pokedex
```go
	// Get Pokedex
	pokedex, err := gokemon.GetPokedex()
	if err != nil {
		panic(err)
	}
	// Print first 3 pokemons
	for _, pokemon := range pokedex.Pokemon[:3] {
		fmt.Println(pokemon.Get())
	}
```

### Pokemons
```go
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
```

### Abilities
```go
	// Get Ability
	swarm, err := gokemon.GetAbility(68)
	if err != nil {
		panic(err)
	}
	fmt.Println(swarm)
```

### Types
```go
	// Get Type
	flying, err := gokemon.GetType(3)
	if err != nil {
		panic(err)
	}
	fmt.Println(flying)
```

### Moves
```go
	// Get Move
	dragonTail, err := gokemon.GetMove(525)
	if err != nil {
		panic(err)
	}
	fmt.Println(dragonTail)
```

### Eggs
```go
	// Get Egg
	monster, err := gokemon.GetEgg(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(monster)
```

### Descriptions
```go
	// Get Description
	desc, err := gokemon.GetDescription(200)
	if err != nil {
		panic(err)
	}
	fmt.Println(desc)
```

### Sprites
```go
	// Get Sprite
	sprite, err := gokemon.GetSprite(77)
	if err != nil {
		panic(err)
	}
	fmt.Println(sprite)
```

### Games
```go
	// Get Game
	yellow, err := gokemon.GetGame(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(yellow)
```
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
	for _, ability := range charizard.Abilities {
		fmt.Println(ability)
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
	// Print first 10 pokemons
	for _, pokemon := range pokedex.Pokemons[:5] {
		fmt.Println(pokemon)
	}
```

### Pokemons
```go
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
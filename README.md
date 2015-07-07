# gokemon
A Golang wrapper for [PokeAPI](http://pokeapi.co/)

## Pokedex
```go
func main() {
	pokedex, err := gokemon.NewPokedex() // HL
	if err != nil {
		panic(err)
	}
	for _, pokemon := range pokedex.Pokemons {
		fmt.Println(pokemon)
	}
}
```
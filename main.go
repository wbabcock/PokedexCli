package main

import (
	"time"

	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &pokeapi.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: client,
	}
	startApp(cfg)
}

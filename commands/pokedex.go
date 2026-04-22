package commands

import (
	"fmt"

	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func pokedex(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}

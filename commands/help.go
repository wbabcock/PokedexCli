package commands

import (
	"fmt"

	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func help(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range ToList() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

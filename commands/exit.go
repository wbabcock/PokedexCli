package commands

import (
	"fmt"
	"os"

	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func exit(cfg *pokeapi.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func catch(cfg *pokeapi.Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	resistence := rand.IntN(pokemon.BaseExperience)
	if resistence > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.CaughtPokemon[pokemon.Name] = pokemon

	return nil
}

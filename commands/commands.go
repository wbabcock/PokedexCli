package commands

import "github.com/wbabcock/pokedexcli/internal/pokeapi"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
}

func ToList() map[string]CliCommand {
	return map[string]CliCommand{
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    map_next,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    map_back,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a location",
			Callback:    explore,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Try to catch a pokemon",
			Callback:    catch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "View details about a caught Pokemon",
			Callback:    inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See all the pokemon you've caught",
			Callback:    pokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    help,
		},
	}
}

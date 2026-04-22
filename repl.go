package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wbabcock/pokedexcli/commands"
	"github.com/wbabcock/pokedexcli/internal/pokeapi"
)

func startApp(cfg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		cmdRequest := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		
		cmd, ok := commands.ToList()[cmdRequest]
		if ok {
			err := cmd.Callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Try 'help' for list of commands.")
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

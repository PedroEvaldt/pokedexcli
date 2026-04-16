package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PedroEvaldt/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var pokemon, mapName string
		if len(words) > 1 {
			if commandName == "explore" {
				mapName = words[1]
			}
			if commandName == "catch" {
				pokemon = words[1]
			}
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, mapName, pokemon)
		if err != nil {
			fmt.Println(err)
		}
		continue
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List the pokemons in the given region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	return words
}

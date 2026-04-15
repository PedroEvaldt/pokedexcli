package repl

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, locName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, locName string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

var (
	nextURL *string
	prevURL *string
)

func commandMap(cfg *Config, locName string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, locName string) error {
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("You are on the first page")
	}

	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandExplore(cfg *Config, locName string) error {
	locationResp, err := cfg.PokeapiClient.SearchLocation(locName)
	if err != nil {
		return err
	}

	var pokemons []string
	for _, pokemonData := range locationResp.PokemonEncounters {
		pokemons = append(pokemons, pokemonData.Pokemon.Name)
	}
	for _, pokemon := range pokemons {
		fmt.Println(pokemon)
	}
	return nil
}

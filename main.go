package main

import (
	"time"

	"github.com/PedroEvaldt/pokedexcli/internal/pokeapi"
	"github.com/PedroEvaldt/pokedexcli/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &repl.Config{
		PokeapiClient: pokeClient,
	}
	repl.StartRepl(cfg)
}

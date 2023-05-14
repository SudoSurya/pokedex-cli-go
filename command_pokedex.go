package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("Your pokedex is empty")
	}
	fmt.Println("Pokemon in Your Pokedex: ")
	for idx, _ := range cfg.caughtPokemon {
		fmt.Println("	-- ", idx)
	}

	return nil

}

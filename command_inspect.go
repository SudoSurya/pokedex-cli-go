package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	resp, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught pokemon yet")
	}
	fmt.Printf("Name: %s\n", resp.Name)
	fmt.Printf("Height: %d\n", resp.Height)
	fmt.Printf("Weight: %d\n", resp.Weight)

	fmt.Println("-- Pokemon Stats: ")
	for _, stat := range resp.Stats {
		fmt.Printf("	--%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("-- Pokemon Types:")
	for _, stat := range resp.Types {
		fmt.Printf("	--%s\n", stat.Type.Name)
	}

	return nil

}

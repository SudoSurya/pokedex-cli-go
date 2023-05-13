package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExpore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationAres(locationAreaName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("pokemons in area: ", resp.Name)
	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" -- %s\n", area.Pokemon.Name)
	}

	return nil

}

package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		log.Fatal(err)
	}

	const threshold = 50

	randomNum := rand.Intn(resp.BaseExperience)
	fmt.Println(resp.BaseExperience, randomNum, threshold)
	if randomNum < threshold {
		return fmt.Errorf("failed to catch the pokemon %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = resp

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil

}

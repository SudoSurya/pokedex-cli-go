package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	_, ok := cfg.caughtPokemon[pokemonName]
	if ok {
		return fmt.Errorf("pokemon already caught")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		log.Fatal(err)
	}

	const threshold = 50

	// randomNum := rand.Intn(resp.BaseExperience)
	// randomNum := rand.Int()
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time
	randNumber := rand.Intn(21) + 40
	fmt.Println(randNumber, threshold)
	if randNumber > threshold {
		return fmt.Errorf("failed to catch the pokemon %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = resp

	fmt.Printf("%s was caught!\n", pokemonName)

	return nil

}

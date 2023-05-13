package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")
	for _, area := range resp.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	return nil

}

func callbackMapb(cfg *config) error {
	if cfg.nextLocURL == nil {
		return errors.New("you're on first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")
	for _, area := range resp.Results {
		fmt.Printf(" -- %s\n", area.Name)
	}
	cfg.nextLocURL = resp.Next
	cfg.prevLocURL = resp.Previous

	return nil

}

package main

import (
	"fmt"
	"log"

	"github.com/20pa5a1210/pokedex-cli-go/internal/pokeapi"
)

func callbackMap() error {

	pokeApiClient := pokeapi.NewClient()
	resp, err := pokeApiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location Areas")
	for _, area := range resp.Results {
		fmt.Printf(" --%s\n", area.Name)
	}

	return nil

}

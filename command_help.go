package main

import "fmt"

func callbackHelp() error{
	fmt.Println("Welcome to Pokedex CLI: ")
	fmt.Println("Here are available Commands: ")
	fmt.Println("")
	commands := getCommands()
	for _, curr := range commands {
		fmt.Printf("	--%s : '%s'\n", curr.name, curr.description)
	}
	fmt.Println("")
	return nil
}

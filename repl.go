package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		command := cleaned[0]

		availableCommands := getCommands()

		commandName, ok := availableCommands[command]
		if !ok {
			fmt.Println("Invalid Command!")
			continue
		}
		err := commandName.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}

}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "prints the help menu",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List some location ares",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List some location ares (map backwards)",
			callback:    callbackMapb,
		},
		"exit": {
			name:        "exit",
			description: "exists from the cli",
			callback:    callbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowStr := strings.ToLower(str)
	words := strings.Fields(lowStr)
	return words
}

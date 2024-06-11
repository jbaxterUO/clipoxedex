package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	command     string
	description string
	callback    func(*Config, ...string) error
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Pokedex CLI!\n")
	for {
		fmt.Print("poxedex> ")
		scanner.Scan()
		text := scanner.Text()

		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		optionalArgs := []string{}

		if len(words) > 1 {
			optionalArgs = words[1:]
		}

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, optionalArgs...)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", commandName)
			fmt.Println("Type 'help' to see the list of available commands")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			command:     "help",
			description: "Prints the list of available commands",
			callback:    callbackHelp,
		},
		"map": {
			command:     "map",
			description: "Displays the next page of areas in the world",
			callback:    callbackMap,
		},
		"mapb": {
			command:     "mapb",
			description: "Displays the previous page of areas in the world",
			callback:    callbackMapB,
		},
		"explore": {
			command:     "explore <location area name>",
			description: "Displays the list of Pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			command:     "catch <pokemon name>",
			description: "Attempts to catch a Pokemon and add it to the Pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			command:     "inspect <pokemon name>",
			description: "Displays the details of a caught Pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			command:     "pokedex",
			description: "Displays the list of caught Pokemon",
			callback:    callbackPokedex,
		},
		"exit": {
			command:     "exit",
			description: "Exits the program",
			callback:    callbackExit,
		},
	}
}

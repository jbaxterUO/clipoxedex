package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("missing argument, please provide a location area name")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Location Area:", locationArea.Name)
	fmt.Println("Pokemon Encounters:")

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println("- ", pokemon.Pokemon.Name)
	}

	return nil
}

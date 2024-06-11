package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("missing argument, please provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemmon has not been caught yet, can only inspect caught pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("\nStats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\n%s: %d", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("\n")
	fmt.Printf("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("%s ", t.Type.Name)
	}
	fmt.Printf("\n")

	return nil
}

package main

import (
	"fmt"
)

func callbackPokedex(cfg *Config, args ...string) error {

	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("you have not caught any pokemon yet")
	}

	fmt.Print("Pokemon in Pokedex:\n")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}

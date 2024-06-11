package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("missing argument, please provide a pokemon name")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught\n", pokemon.Name)
	return nil
}

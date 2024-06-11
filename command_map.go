package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *Config, args ...string) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	fmt.Println("Location Areas:")

	for _, locationArea := range resp.Results {
		fmt.Println("- ", locationArea.Name)
	}

	return nil
}

func callbackMapB(cfg *Config, args ...string) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are at the first area, try using the map command to go forward")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	fmt.Println("Location Areas:")

	for _, locationArea := range resp.Results {
		fmt.Println("- ", locationArea.Name)
	}

	return nil
}

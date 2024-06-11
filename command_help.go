package main

import (
	"fmt"
)

func callbackHelp(cfg *Config, args ...string) error {
	fmt.Println("Usage:")
	for _, command := range getCommands() {
		fmt.Printf("  %s: %s\n", command.command, command.description)
	}
	return nil
}

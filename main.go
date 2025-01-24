package main

import (
	"fmt"
	_ "gocli/cmd"
	"gocli/core"
	"os"
)

func main() {
	commandName := os.Args[1]

	if command, exists := core.Commands[commandName]; exists {
		if err := command.Execute(os.Args[2:]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Unknown command: %s\n", commandName)
		fmt.Println("Available commands: ")
		for _, cmd := range core.Commands {
			fmt.Printf("  %s: %s\n", cmd.Name, cmd.Description)
		}
		os.Exit(1)
	}
}

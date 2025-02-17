package main

import (
	"fmt"
	_ "gocli/cmd"
	"gocli/core"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		showUsage()
		log.Fatalf("Issue")
	}

	cmdName := os.Args[1]
	cmd, exists := core.Commands[cmdName]
	if !exists {
		fmt.Printf("Unknown command: %s\n\n", cmdName)
		showUsage()
		os.Exit(1)
	}

	err := cmd.Execute(os.Args[2:])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("Usage: mycli <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	for name, cmd := range core.Commands {
		fmt.Printf("  %-10s %s\n", name, cmd.Help())
	}
}

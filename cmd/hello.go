package cmd

import (
	"fmt"
	"gocli/core"
)

func init() {
	core.RegisterCommand(
		"hello",
		"This commands prints Hello, World!",
		Hello,
		HelloHelp)
}

func Hello(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("hello() expects no arguments, but got: %d", len(args))
	}

	fmt.Println("Hello, World!")
	return nil
}

func HelloHelp() {
	fmt.Println("Usage: hello")
	fmt.Println("Sanity check to make sure gocli is loaded in terminal.")
}

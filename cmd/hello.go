package cmd

import (
	"fmt"
	"gocli/core"
)

type HelloCommand struct {
	Name        string
	Description string
}

func init() {
	core.Commands["hello"] = &HelloCommand{
		Name:        "hello",
		Description: "prints Hello, World!",
	}
}

func (h *HelloCommand) Execute(args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("hello() expects no arguments, but got: %d", len(args))
	}

	fmt.Println("Hello, World!")
	return nil
}

func (h *HelloCommand) Help() string {
	fmt.Println("Usage: hello")
	fmt.Println("Sanity check to make sure gocli is loaded in terminal.")
	return h.Description
}

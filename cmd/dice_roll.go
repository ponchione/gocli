package cmd

import (
	"fmt"
	"gocli/core"
	"gocli/util"
	"math/rand"
	"strconv"
)

type RollCommand struct {
	Name        string
	Description string
}

func init() {
	core.Commands["roll"] = &RollCommand{
		Name:        "roll",
		Description: "Generates a random number between the specified range.",
	}
}

func (r *RollCommand) Execute(args []string) error {
	if err := validateNumberArgs(args); err != nil {
		return err
	}

	mini, err1 := strconv.Atoi(args[0])
	maxi, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		return fmt.Errorf("couldn't convert input(s). Please enter valid integers")
	}

	if mini > maxi {
		return fmt.Errorf("min value cannot be greater than max value")
	}

	random := rand.Intn(maxi-mini+1) + mini
	fmt.Printf("roll = %d", random)

	return nil
}

func (r *RollCommand) Help() string {
	fmt.Println("Usage: mycli roll <min> <max>")
	fmt.Println("Example: mycli roll 1 6")
	fmt.Println("Generates a random number between the specified range.")

	return r.Description
}

func validateNumberArgs(args []string) error {
	if err := util.ValidateArgs(args, ""); err != nil {
		return err
	}

	if len(args) < 2 {
		return fmt.Errorf("roll requires at least 2 integers")
	}

	if len(args) > 2 {
		return fmt.Errorf("roll only takes 2 integer values")
	}

	return nil
}

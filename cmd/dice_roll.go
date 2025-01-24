package cmd

import (
	"fmt"
	"gocli/core"
	"gocli/util"
	"math/rand"
	"strconv"
)

func init() {
	core.RegisterCommand("roll",
		"roll command returns a random number between the user specified range (must be integers)",
		Roll,
		RollHelp)
}

func Roll(args []string) error {
	if err := ValidateNumberArgs(args); err != nil {
		return err
	}

	min, err1 := strconv.Atoi(args[0])
	max, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		return fmt.Errorf("couldn't convert input(s). Please enter valid integers")
	}

	if min > max {
		return fmt.Errorf("min value cannot be greater than max value")
	}

	random := rand.Intn(max-min+1) + min
	fmt.Printf("roll = %d", random)

	return nil
}

func RollHelp() {
	fmt.Println("Usage: mycli roll <min> <max>")
	fmt.Println("Example: mycli roll 1 6")
	fmt.Println("Generates a random number between the specified range.")
}

func ValidateNumberArgs(args []string) error {
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

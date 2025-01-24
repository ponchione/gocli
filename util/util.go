package util

import (
	"fmt"
)

func ValidateArgs(args []string, msg string) error {
	if msg == "" {
		msg = "No commands issued to gocli"
	}

	if len(args) == 0 {
		return fmt.Errorf("%s", msg)
	}

	return nil
}

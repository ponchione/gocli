package util

import (
	"fmt"
)

func ValidateArgs(args []string, msg string) error {
	//TODO this won't catch no commands, this is only for args for commands. Change this.
	if msg == "" {
		msg = "No commands issued to gocli"
	}

	if len(args) == 0 {
		return fmt.Errorf("%s", msg)
	}

	return nil
}

package cmd

import (
	"fmt"
	"gocli/core"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	core.RegisterCommand(
		"ex",
		"command to open the file explore at specified location",
		Explorer,
		ExplorerHelp)
}

func Explorer(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("ex only accepts one possible command: path (optional)")
	}

	var path string
	if len(args) == 1 {
		path = args[0]
	} else {
		var err error
		path, err = os.Getwd()
		if err != nil {
			log.Fatalf("error getting the current directory: %v", err)
		}
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Invalid path: %v", err)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		log.Fatalf("Path does not exist: %s\n", info)
	} else if err != nil {
		log.Fatalf("error checing path: %s\n", info)
	}

	if !info.IsDir() {
		log.Fatalf("Path is not a directory: %s\n", info)
	}

	cmd := exec.Command("explorer.exe", absPath)
	err = cmd.Start()
	if err != nil {
		log.Fatalf("Error opening file explorer: %v", err)
	}

	return nil
}

func ExplorerHelp() {
	fmt.Println("Usage: ex <path>(optional)")
	fmt.Println("Example: ex . will open the file explorer at current directory level")
	fmt.Println("command to open the file explore at specified location")
}

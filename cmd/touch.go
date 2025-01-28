package cmd

import (
	"fmt"
	"gocli/core"
	"gocli/util"
	"log"
	"os"
	"path/filepath"
)

func init() {
	core.RegisterCommand(
		"touch",
		"touch command creates a file",
		Touch,
		TouchHelp)
}

func Touch(args []string) error {
	if err := util.ValidateArgs(args, "touch requires at least one argument."); err != nil {
		return err
	}

	if len(args) > 2 {
		return fmt.Errorf("touch only take two arguments: file name | path(optional)")
	}

	fileName := args[0]
	var dir string

	if len(args) == 2 {
		dir = args[1]
		info, err := os.Stat(dir)

		if os.IsNotExist(err) {
			return fmt.Errorf("provided file path does not exist: %s", dir)
		}

		if !info.IsDir() {
			return fmt.Errorf("provided file path is not a directory: %s", dir)
		}
	} else {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %v", err)
		}
	}

	if filepath.Ext(fileName) == "" {
		fileName += ".txt"
	}

	filePath := filepath.Join(dir, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("error occured when closing file. Error: %v", err)
		}
	}(file)

	fmt.Printf("File created at: %s\n", filePath)
	return nil
}

func TouchHelp() {
	fmt.Println("Usage: touch <file> <path>(optional)")
	fmt.Println("Example: touch myText.txt will create a text file called myText at the current directory")
	fmt.Println("touch is used to create a file via command line.")
}

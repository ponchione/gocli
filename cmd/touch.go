package cmd

import (
	"fmt"
	"gocli/core"
	"log"
	"os"
	"path/filepath"
)

type TouchCommand struct {
	Name        string
	Description string
}

func init() {
	core.Commands["touch"] = &TouchCommand{
		Name:        "touch",
		Description: "touch is used to create a file via command line.",
	}
}

func (t *TouchCommand) Execute(args []string) error {
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
		err = file.Close()
		if err != nil {
			log.Fatalf("error occured when closing file. Error: %v", err)
		}
	}(file)

	fmt.Printf("File created at: %s\n", filePath)
	return nil
}

func (t *TouchCommand) Help() string {
	fmt.Println("Usage: touch <file> <path>(optional)")
	fmt.Println("Example: touch myText.txt will create a text file called myText at the current directory")
	fmt.Println("touch is used to create a file via command line.")
	return t.Description
}

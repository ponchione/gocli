package cmd

import (
	"fmt"
	"gocli/core"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OrganizeCommand struct {
	Name        string
	Description string
}

func init() {
	core.Commands["organize"] = &OrganizeCommand{
		Name:        "organize",
		Description: "organize will organize the contents of your Downloads folder.",
	}
}

var extensions = map[string]bool{
	".zip":  true,
	".png":  true,
	".jpeg": true,
	".jpg":  true,
	".msi":  true,
	".txt":  true,
	".pdf":  true,
}

func (o *OrganizeCommand) Execute(args []string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	downloadsPath := filepath.Join(homeDir, "Downloads")

	contents, err := os.ReadDir(downloadsPath)
	if err != nil {
		log.Fatalf("Error reading contents of Downloads directory: %v", err)
	}

	for _, file := range contents {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if !extensions[ext] {
			continue
		}

		extName := ext[1:]
		targetDir := filepath.Join(downloadsPath, extName+"_files")
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			log.Printf("error creating directory for %s: %v\n", extName, err)
			continue
		}

		sourcePath := filepath.Join(downloadsPath, file.Name())
		destPath := filepath.Join(targetDir, file.Name())
		if err := moveFile(sourcePath, destPath); err != nil {
			log.Printf("error moving file %s: %v", file.Name(), err)
			continue
		}

		fmt.Printf("Moved %s to %s\n", file.Name(), filepath.Base(targetDir))

	}

	fmt.Println("Sorting completed.")

	return nil
}

func moveFile(sourcePath string, destPath string) error {
	if _, err := os.Stat(destPath); err == nil {
		baseName := filepath.Base(destPath)
		ext := filepath.Ext(baseName)
		nameWithoutExt := strings.TrimSuffix(baseName, ext)

		counter := 1
		for {
			newName := fmt.Sprintf("%s_%d%s", nameWithoutExt, counter, ext)
			destPath := filepath.Join(filepath.Dir(destPath), newName)

			if _, err := os.Stat(destPath); os.IsNotExist(err) {
				break
			}
			counter++
		}
	}

	if err := os.Rename(sourcePath, destPath); err == nil {
		return nil
	}

	return nil
}

func (o *OrganizeCommand) Help() string {
	return o.Description
}

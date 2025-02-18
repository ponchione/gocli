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

var extensions = map[string]string{
	".zip":  "zip_files",
	".png":  "image_files",
	".jpeg": "image_files",
	".jpg":  "image_files",
	".msi":  "installers",
	".txt":  "text_files",
	".pdf":  "pdf_files",
}

var folderNames = []string{
	"zip_files",
	"image_files",
	"pdf_files",
	"old_files",
	"installers",
	"csv_files",
	"spreadsheets",
	"text_files",
}

func (o *OrganizeCommand) Execute(args []string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}

	downloadsPath := filepath.Join(homeDir, "Downloads")
	createFolders(downloadsPath)

	contents, err := os.ReadDir(downloadsPath)
	if err != nil {
		log.Fatalf("Error reading contents of Downloads directory: %v", err)
	}

	for _, file := range contents {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if _, exists := extensions[ext]; !exists {
			log.Printf("extension for %s does not exist for sorting\n", file)
			continue
		}

		targetDir := filepath.Join(downloadsPath, extensions[ext])
		sourcePath := filepath.Join(downloadsPath, file.Name())
		destPath := filepath.Join(targetDir, file.Name())
		if err := moveFile(sourcePath, destPath); err != nil {
			log.Printf("error moving file %s: %v", file.Name(), err)
			continue
		}
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

func createFolders(dPath string) {
	for _, name := range folderNames {
		target := filepath.Join(dPath, name)
		err := os.MkdirAll(target, 0755)
		if err != nil {
			log.Printf("error creating directory for %s: %v\n", name, err)
			continue
		} else {
			log.Printf("folder %s already exists\n", name)
		}
	}
}

func (o *OrganizeCommand) Help() string {
	return o.Description
}

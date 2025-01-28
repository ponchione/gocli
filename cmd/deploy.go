package cmd

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/joho/godotenv"
	"gocli/core"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func init() {
	core.RegisterCommand("deploy",
		"Deploys latest changes made to gocli",
		Deploy,
		DeployHelp)
}

const exeName = "gocli.exe"

func Deploy(args []string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	tempDir, err := os.MkdirTemp("", "temp-deploy-*")
	if err != nil {
		return fmt.Errorf("failed to create temporary directory: %v", err)
	}
	defer func(tempDir string) {
		if err := os.RemoveAll(tempDir); err != nil {
			log.Printf("Error occured when trying to remove %s: %v", tempDir, err)
		}
	}(tempDir)
	log.Printf("Temporary directory created at: %s\n", tempDir)

	log.Printf("Attempting to clone github repo...")
	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:      os.Getenv("CLI_GITHUB"),
		Progress: os.Stdout,
	})
	if err != nil {
		return fmt.Errorf("failed to clone git repository: %v", err)
	}

	log.Println("Building gocli project...")
	buildCmd := exec.Command("go", "build", "-o", exeName)
	buildCmd.Dir = tempDir
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("failed to build project: %v", err)
	}

	outDir := os.Getenv("DEPLOY_DIR")
	log.Printf("Moving gocli.exe to %s\n", outDir)
	from := filepath.Join(tempDir, exeName)
	to := filepath.Join(outDir, exeName)
	if err := os.Rename(from, to); err != nil {
		return fmt.Errorf("failed to move %s from %s to %s : %v",
			exeName, from, to, err)
	}

	return nil
}

func DeployHelp() {
	fmt.Println("Usage: deploy")
}

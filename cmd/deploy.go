package cmd

/*
TODO:
This program fails to work due to Windows locking the existing gocli.exe
(because it's in use from Powershell), so this command can't remove and
then replace it.

Need to rethink approach here.
*/

//func init() {
//	core.RegisterCommand(
//		"deploy",
//		"Deploys latest changes made to gocli",
//		Deploy,
//		DeployHelp,
//	)
//}
//
//const (
//	exeName = "gocli.exe"
//)
//
//func Deploy(args []string) error {
//	Validate(args) //Make sure there are none
//
//	//Create temp work directory
//	tempDir, err := os.MkdirTemp("", "temp-deploy-*")
//	if err != nil {
//		return fmt.Errorf("failed to create temporary directory: %v", err)
//	}
//	defer func(tempDir string) {
//		if err := os.RemoveAll(tempDir); err != nil {
//			log.Printf("Error occured when trying to remove %s: %v", tempDir, err)
//		}
//	}(tempDir)
//	log.Printf("Temporary directory created at: %s\n", tempDir)
//
//	//Clone the repo
//	log.Printf("Attempting to clone github repo...")
//	githubUrl := "https://github.com/ponchione/gocli.git"
//	_, err = git.PlainClone(tempDir, false, &git.CloneOptions{
//		URL:      githubUrl,
//		Progress: os.Stdout,
//	})
//	if err != nil {
//		return fmt.Errorf("failed to clone git repository: %v", err)
//	}
//
//	//Copy local .env file
//	log.Println("Copying .env file...")
//	srcEnvPath := "C:\\Users\\mitch\\source\\bin\\.env"
//	srcEnv, err := os.Open(srcEnvPath)
//	if err != nil {
//		log.Fatalf("couldn't open local .env file: %v", err)
//	}
//	defer func(srcEnv *os.File) {
//		if err := srcEnv.Close(); err != nil {
//			log.Fatalf("error occured when closing file. Error: %v", err)
//		}
//	}(srcEnv)
//
//	destEnvPath := filepath.Join(tempDir, ".env")
//	destEnv, err := os.Create(destEnvPath)
//	if err != nil {
//		log.Fatalf("error creating .env at destination: %v", err)
//	}
//	defer func(destEnv *os.File) {
//		if err := destEnv.Close(); err != nil {
//			log.Fatalf("error closing new .env file: %v", err)
//		}
//	}(destEnv)
//
//	if _, err := io.Copy(destEnv, srcEnv); err != nil {
//		log.Fatalf("error copying contents from old .env to new .env: %v", err)
//	}
//
//	//Load .env file
//	loadPath := filepath.Join(tempDir, ".env")
//	err = godotenv.Load(loadPath)
//	if err != nil {
//		return fmt.Errorf("error loading .env file: %v", err)
//	}
//
//	//Build project
//	log.Println("Building gocli project...")
//	buildCmd := exec.Command("go", "build", "-o", exeName)
//	buildCmd.Dir = tempDir
//	buildCmd.Stdout = os.Stdout
//	buildCmd.Stderr = os.Stderr
//	if err := buildCmd.Run(); err != nil {
//		return fmt.Errorf("failed to build project: %v", err)
//	}
//
//	//Remove old gocli.exe
//	outDir := "C:\\Users\\mitch\\source\\bin\\"
//	if err := os.Remove(outDir + "gocli.exe"); err != nil {
//		return fmt.Errorf("failed to remove old gocli.exe: %v", err)
//	}
//
//	//"Deploy" project to bin
//	log.Printf("Moving gocli.exe to %s...\n", outDir)
//	from := filepath.Join(tempDir, exeName)
//	to := filepath.Join(outDir, exeName)
//	if err := os.Rename(from, to); err != nil {
//		return fmt.Errorf("failed to move %s from %s to %s : %v",
//			exeName, from, to, err)
//	}
//
//	log.Println("Successfully deployed new version of gocli")
//
//	return nil
//}
//
//func Validate(args []string) {
//	if len(args) > 0 {
//		log.Println("error: deploy does not take any arguments")
//		os.Exit(1)
//	}
//}
//
//func DeployHelp() {
//	fmt.Println("Usage: deploy")
//}

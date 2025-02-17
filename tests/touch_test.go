package tests

//func TestTouch_CreateFileWithDefaultExtension(t *testing.T) {
//	tempDir := os.TempDir()
//	fileName := "testfile"
//	filePath := filepath.Join(tempDir, fileName+".txt")
//
//	err := cmd.Touch([]string{fileName, tempDir})
//	if err != nil {
//		t.Fatalf("Expected no error, but got: %v", err)
//	}
//
//	if _, err := os.Stat(filePath); os.IsNotExist(err) {
//		t.Fatalf("Expected file to be created at %s", filePath)
//	}
//	defer os.Remove(filePath)
//}
//
//func TestTouch_CreateFileWithProvidedExtension(t *testing.T) {
//	tempDir := os.TempDir()
//	fileName := "testfile.log"
//	filePath := filepath.Join(tempDir, fileName)
//
//	err := cmd.Touch([]string{fileName, tempDir})
//	if err != nil {
//		t.Fatalf("Expected no error, but got: %v", err)
//	}
//
//	if _, err := os.Stat(filePath); os.IsNotExist(err) {
//		t.Fatalf("Expected file to be created at %s", filePath)
//	}
//	defer os.Remove(filePath)
//}
//
//func TestTouch_InvalidDirectory(t *testing.T) {
//	err := cmd.Touch([]string{"testfile", "/invalid/path"})
//	if err == nil {
//		t.Fatalf("Expected an error due to invalid directory, but got none")
//	}
//}
//
//func TestTouch_NoArgumentsProvided(t *testing.T) {
//	err := cmd.Touch([]string{})
//	if err == nil {
//		t.Fatalf("Expected an error due to missing arguments, but got none")
//	}
//}

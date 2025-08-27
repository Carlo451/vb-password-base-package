package tests

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Carlo451/vb-password-base-package/api"
	"github.com/Carlo451/vb-password-base-package/environment"
)

const (
	basePath  = "/home/carl-moritz/vB-Password-Store"
	storeName = "testStore"
)

func setup() api.PasswordStoreHandler {
	os.Mkdir(basePath, 0755)
	os.Setenv("VB_PASSWORDSTORE_BASE_DIR", basePath)
	path, _ := environment.LookUpAndGetEnvValue("VB_PASSWORDSTORE_BASE_DIR")
	handler := api.NewPasswordStoreHandler(path)
	handler.CreatePasswordStore("testStore", "camo", "Id")
	return handler
}
func teardown() {
	err := os.RemoveAll(basePath)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreatePasswordStore(t *testing.T) {
	setup()
	_, err := os.ReadDir(filepath.Join(basePath, storeName+"/configs"))
	if err != nil {
		log.Fatal(err)
		t.Error("Expected Password Store to be created")
	}
	teardown()
}

func TestAddContentDirInRootStore(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName), storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestAddContentDirInRootStore_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("", storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestAddContentDirInRootStoreWitMultipleNonExistingSubDirs(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/doubleextraContent", storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestAddContentDirInRootStoreWitMultipleExistingSubDirs(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/doubleextraContent", storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore("/extraContent", storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/extraContent/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestInsertContentIntoContentDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "camo123", "username")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

func TestInsertContentIntoContentDir_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.InsertContentInContentDirectory("/extraContent/content", storeName, "camo123", "username")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

func TestInsertContentWhenIdentifierIsAlreadyUsed(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	_, err := handler.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "camo123", "password")
	if err == nil {
		t.Errorf("This insertion should fail because identifier is already used")
	}

	teardown()
}

func TestInsertContentWhenContentDirDoesNotExist(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	_, err := handler.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/NotExistingCOntent"), storeName, "camo123", "password")
	if err == nil {
		t.Errorf("This insertion should fail because the content directory does not exist")
	}
	teardown()
}

func TestUpdateContentInContentDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.UpdateContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "camo123", "password")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

func TestUpdateContentInContentDir_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/doubleextraContent", storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore("/extraContent", storeName, "content", "password123", "password")
	check, _ := handler.UpdateContentInContentDirectory("/extraContent/content", storeName, "camo123", "password")
	if !check {
		t.Errorf("Content should have been updated")
	}
	teardown()
}

func TestUpdateContentInContentDirWithNoCorrespondingContentFile(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.UpdateContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "password", "username")
	if !check {
		t.Errorf("Content should have been updated")
	}
	teardown()
}

func TestRemoveDirectory(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.RemoveDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, true)
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	if !api.CheckIfDirectoryExists(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent")) {
		t.Errorf("removeSubdirs is true but in extraContent is still another dir so this should still exists")
	}
	teardown()
}

func TestRemoveDirectory_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.RemoveDirectory("/extraContent/content", storeName, true)
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	if !api.CheckIfDirectoryExists(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent")) {
		t.Errorf("removeSubdirs is true but in extraContent is still another dir so this should still exists")
	}
	teardown()
}

func TestRemoveDirectoryWithDeletionOfEmptySubDirs(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.RemoveDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, true)
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	if api.CheckIfDirectoryExists(filepath.Join(basePath, storeName+"/extraContent")) {
		t.Errorf("removeSubdirs is true and content was delted so this dir should be removed aswell")
	}
	teardown()
}

func TestRemoveDirectoryWithoutDeletionOfEmptySubDirs(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	check, _ := handler.RemoveDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, false)
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	if !api.CheckIfDirectoryExists(filepath.Join(basePath, storeName+"/extraContent")) {
		t.Errorf("removeSubdirs is false so even the the subDir is empty it should not be deleted")
	}
	teardown()
}

func TestReadContentDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	contentDir, err := handler.ReadContentDir(filepath.Join(basePath, storeName+"/extraContent/content"), storeName)
	if err != nil {
		t.Errorf("COntentDirectory should not return an error")
	}
	for _, file := range contentDir.ReturnFiles() {
		if file.GetFileName() == "password" {
			password := file.GetContent()
			if password != "password123" {
				t.Errorf("Something went wrong")
			}
		}
	}
	teardown()
}

func TestRemoveContentInContentDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	handler.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "password123", "username")
	error := handler.DeleteContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "password")
	if error != nil {
		t.Errorf("Something went wrong when deleting content")
	}
	if api.CheckIfContentFileExists(filepath.Join(basePath, storeName, "/extraContent/content/password")) {
		t.Errorf("This File should have been removed")
	}
	if !api.CheckIfContentFileExists(filepath.Join(basePath, storeName, "/extraContent/content/username")) {
		t.Errorf("This File should still be in the directory")
	}

	teardown()
}
func TestRemoveContentInContentDir_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent", storeName, "content", "password123", "username")
	handler.InsertContentInContentDirectory("/extraContent/content", storeName, "password123", "password")
	error := handler.DeleteContentInContentDirectory("/extraContent/content", storeName, "password")
	if error != nil {
		t.Errorf("Something went wrong when deleting content")
	}
	if api.CheckIfContentFileExists(filepath.Join(basePath, storeName, "/extraContent/content/password")) {
		t.Errorf("This File should have been removed")
	}
	if !api.CheckIfContentFileExists(filepath.Join(basePath, storeName, "/extraContent/content/username")) {
		t.Errorf("This File should still be in the directory")
	}
	teardown()
}

func TestRemoveContentInContentDir_AndRemoveEmptyContentDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), storeName, "content", "password123", "password")
	error := handler.DeleteContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), storeName, "password")
	if error != nil {
		t.Errorf("Something went wrong when deleting content")
	}
	contentDir, _ := api.CheckIfContentDirectoryExists(filepath.Join(basePath, storeName, "/extraContent/content"))
	if contentDir {
		t.Errorf("This contnetDIr was Empty and should have been removed too")
	}
	store := api.CheckIfDirectoryExists(filepath.Join(basePath, storeName, "/extraContent"))
	if store {
		t.Errorf("Since extra content directory is empty too after the content dir was deleted, it should have been removed too.")
	}
	teardown()
}

func TestRemoveContentInContentDir_AndRemoveEmptyContentDir_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent", storeName, "content", "password123", "password")
	error := handler.DeleteContentInContentDirectory("/extraContent/content", storeName, "password")
	if error != nil {
		t.Errorf("Something went wrong when deleting content")
	}
	contentDir, _ := api.CheckIfContentDirectoryExists(filepath.Join(basePath, storeName, "/extraContent/content"))
	if contentDir {
		t.Errorf("This contnetDIr was Empty and should have been removed too")
	}
	store := api.CheckIfDirectoryExists(filepath.Join(basePath, storeName, "/extraContent"))
	if store {
		t.Errorf("Since extra content directory is empty too after the content dir was deleted, it should have been removed too.")
	}
	teardown()
}

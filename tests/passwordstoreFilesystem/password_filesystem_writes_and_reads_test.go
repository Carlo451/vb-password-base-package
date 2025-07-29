package passwordstoreFilesystem

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"vb-password-store-base/api"
	"vb-password-store-base/environment"
)

const (
	basePath  = "/home/carl-moritz/vB-Password-Store"
	storeName = "testStore"
)

func setup() {
	os.Mkdir(basePath, 0755)
	os.Setenv("VB_PASSWORDSTORE_BASE_DIR", basePath)
	path, _ := environment.LookUpAndGetEnvValue("VB_PASSWORDSTORE_BASE_DIR")
	api.CreatePasswordStore(path, "testStore", "camo", "Id")
}
func teardown() {
	err := os.RemoveAll(basePath)
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreatePasswordStore(t *testing.T) {
	setup()
	_, err := os.ReadDir(filepath.Join(basePath, storeName))
	if err != nil {
		log.Fatal(err)
		t.Error("Expected Password Store to be created")
	}
	teardown()
}

func TestAddContentDirInRootStore(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName), basePath, storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestAddContentDirInRootStoreWitMultipleNonExistingSubDirs(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestAddContentDirInRootStoreWitMultipleExistingSubDirs(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	check := api.CheckIfContentFileExists(filepath.Join(basePath, storeName+"/extraContent/content/password"))
	if !check {
		t.Error("Expected content file to exist")
	}
	teardown()
}

func TestInsertContentIntoContentDir(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	check, _ := api.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), basePath, storeName, "camo123", "username")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

func TestInsertContentWhenIdentifierIsAlreadyUsed(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	_, err := api.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), basePath, storeName, "camo123", "password")
	if err == nil {
		t.Errorf("This insertion should fail because identifier is already used")
	}
	teardown()
}

func TestInsertContentWhenContentDirDoesNotExist(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	_, err := api.InsertContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/NotExistingCOntent"), basePath, storeName, "camo123", "password")
	if err == nil {
		t.Errorf("This insertion should fail because the content directory does not exist")
	}
	teardown()
}

func TestUpdateContentInContentDir(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	check, _ := api.UpdateContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), basePath, storeName, "camo123", "password")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

func TestUpdateContentInContentDirWithNoCorrespondingContentFile(t *testing.T) {
	setup()
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent/doubleextraContent"), basePath, storeName, "content", "password123", "password")
	api.AddContentDirectoryToStore(filepath.Join(basePath, storeName+"/extraContent"), basePath, storeName, "content", "password123", "password")
	check, _ := api.UpdateContentInContentDirectory(filepath.Join(basePath, storeName+"/extraContent/content"), basePath, storeName, "password", "username")
	if !check {
		t.Errorf("Inserted content should have been inserted")
	}
	teardown()
}

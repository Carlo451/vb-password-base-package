package api

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"vb-password-store-base/passwordstore/passwordstoreFilesystem"
	"vb-password-store-base/pathparser"
)

func CreatePasswordStore(path, name, user, encryptionId string) passwordstoreFilesystem.PasswordStoreDir {
	return CreateRootDir(path, name, user, encryptionId)
}

func ReadPasswordStore(path, name string) passwordstoreFilesystem.PasswordStoreDir {
	var rootdir passwordstoreFilesystem.PasswordStoreDir = *CreateNotLoadedRootDir(path, name)
	rootdir.ReadDirectory()
	return rootdir
}

// AddContentDirectoryToStore Adds a new Content directory with the given content and identifiers as well as the corresponding subdirectories in the path
// path - path to the last sub dir, where the new contentDir should be created
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// contentDirName - the name of the content directory
// content - the content
// identifier - the name of the file
func AddContentDirectoryToStore(path, basePath, storeName, contentDirName, content, identifier string) {
	parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(basePath, storeName), filepath.Join(path, contentDirName))
	store := ReadPasswordStore(basePath, storeName)
	dir, exists, remaining := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
	if exists {
		contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(dir, parsedPath.ContentDirectory)
		addAndWriteContentToContentDirectory(content, identifier, contentDir)
	} else {
		lastCreatedSubDir := createSubDirectoriesRek(dir, remaining)
		contentDir := passwordstoreFilesystem.NewEmptyContentDirecotry(lastCreatedSubDir, parsedPath.ContentDirectory)
		addAndWriteContentToContentDirectory(content, identifier, contentDir)
	}
}

// InsertContentInContentDirectory inserts a new contentFile into an already existing content Directory
// path - path to the content Directory
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// content - the content
// identifier - the name of the file
func InsertContentInContentDirectory(path, basePath, storeName, content, identifier string) (bool, error) {
	dirExists, err := CheckIfContentDirectoryExists(path)
	if err != nil {
		return false, err
	}
	if dirExists {
		if CheckIfContentFileExists(filepath.Join(path, identifier)) {
			return false, errors.New("content file already exists, it needs to be updated not inserted")
		}
		store := ReadPasswordStore(basePath, storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(basePath, storeName), path)
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			contentDirs := dir.GetContentDirectories()
			for _, contentDir := range contentDirs {
				if contentDir.GetDirName() == parsedPath.ContentDirectory {
					contentDir.CreateAndAppend(content, identifier)
					return true, nil
				}
			}
		}
	}
	return false, errors.New("the content directory in the Path does not exist")
}

func CheckIfContentDirectoryExists(path string) (bool, error) {
	entry, err := os.ReadDir(path)
	if err != nil {
		return false, nil
	}
	for _, entry := range entry {
		if entry.IsDir() {
			return false, errors.New("the directory exists, but is not a content Directory")
		} else {
			return true, nil
		}
	}
	log.Default().Println("Seems like a empty directory exists, but is not a content Directory -> it will get cleanedUp")
	CleanUpContentDirectory(path)
	return false, errors.New("the directory does not exist")

}

func CleanUpContentDirectory(path string) error {
	return os.RemoveAll(path)
}

func CheckIfContentFileExists(path string) bool {
	_, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	return true
}

package api

import (
	"errors"
	"log"
	"os"
)

func CheckIfDirectoryExists(path string) bool {
	_, err := os.ReadDir(path)
	if err != nil {
		return false
	}
	return true
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

func NewPasswordStoreHandler(path string) PasswordStoreHandler {
	return PasswordStoreHandler{path: path}
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
	log.Default().Println("Seems like a empty directory exists, but is not a sub Directory -> it will get cleanedUp")
	CleanUpContentDirectory(path)
	return false, errors.New("the directory does not exist")
}

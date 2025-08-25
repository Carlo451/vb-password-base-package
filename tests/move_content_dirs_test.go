package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestContentDirectoryMoveToAnotherDirectory(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/Subdir/youtube/personal", storeName, "content", "password123", "password")
	handler.MoveDirectory(filepath.Join(basePath, storeName, "/extraContent/Subdir/youtube/personal/content"), storeName, "/extraContent/movedContent")
	contentDir, err := handler.ReadContentDir("/extraContent/movedContent/content", storeName)
	if err != nil {
		t.Errorf("ContentDirectory should not return an error")
	}
	for _, file := range contentDir.ReturnFiles() {
		if file.GetFileName() == "password" {
			password := file.GetContent()
			if password != "password123" {
				t.Errorf("Something went wrong")
			}
		}
	}
	// empty subdirs should be deleted
	_, errEmptyDir := os.Stat(filepath.Join(basePath, storeName+"/extraContent/Subdir/"))
	if errEmptyDir == nil {
		t.Errorf("dir should not exist anymore")
	}
	teardown()
}

func TestContentDirectoryMoveToAnotherDirectory_RelativePaths(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/Subdir/youtube/personal", storeName, "content", "password123", "password")
	handler.MoveDirectory("/extraContent/Subdir/youtube/personal/content", storeName, "/extraContent/movedContent")
	contentDir, err := handler.ReadContentDir("/extraContent/movedContent/content", storeName)
	if err != nil {
		t.Errorf("ContentDirectory should not return an error")
	}
	for _, file := range contentDir.ReturnFiles() {
		if file.GetFileName() == "password" {
			password := file.GetContent()
			if password != "password123" {
				t.Errorf("Something went wrong")
			}
		}
	}
	// empty subdirs should be deleted
	_, errEmptyDir := os.Stat(filepath.Join(basePath, storeName+"/extraContent/Subdir/"))
	if errEmptyDir == nil {
		t.Errorf("dir should not exist anymore")
	}
	teardown()
}

func TestContentDirectoryMoveToAnotherDirectory_DeletionTillSubDir(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/Subdir/youtube/personal", storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore("/extraContent/Subdir/", storeName, "content", "password123", "password")
	handler.MoveDirectory("/extraContent/Subdir/youtube/personal/content", storeName, "/extraContent/movedContent")
	contentDir, err := handler.ReadContentDir("/extraContent/movedContent/content", storeName)
	if err != nil {
		t.Errorf("ContentDirectory should not return an error")
	}
	for _, file := range contentDir.ReturnFiles() {
		if file.GetFileName() == "password" {
			password := file.GetContent()
			if password != "password123" {
				t.Errorf("Something went wrong")
			}
		}
	}
	// empty subdirs should be deleted
	_, errEmptyDir := os.Stat(filepath.Join(basePath, storeName+"/extraContent/Subdir/youtube"))
	if errEmptyDir == nil {
		t.Errorf("dir should not exist anymore")
	}
	// in Subdir is another contentDir should not be deleted
	_, errNotEmptyDir := os.Stat(filepath.Join(basePath, storeName+"/extraContent/Subdir/"))
	if errNotEmptyDir != nil {
		t.Errorf("dir should not exist anymore")
	}
	teardown()
}

func TestContentDirectoryMoveToAnotherSubDir_MultipleContents(t *testing.T) {
	handler := setup()
	handler.AddContentDirectoryToStore("/extraContent/Subdir/youtube/personal", storeName, "content", "password123", "password")
	handler.AddContentDirectoryToStore("/extraContent/Subdir/youtube/personal", storeName, "anotherContent", "password123", "password")
	handler.InsertContentInContentDirectory("/extraContent/Subdir/youtube/personal/content", storeName, "username", "username")
	handler.MoveDirectory("/extraContent/Subdir/youtube/personal/content", storeName, "/extraContent/movedContent")
	contentDir, err := handler.ReadContentDir("/extraContent/movedContent/content", storeName)
	if err != nil {
		t.Errorf("ContentDirectory should not return an error")
	}
	if len(contentDir.ReturnFiles()) != 2 {
		t.Errorf("contentDir should contain two files")
	}
	for _, file := range contentDir.ReturnFiles() {
		if file.GetFileName() == "password" {
			password := file.GetContent()
			if password != "password123" {
				t.Errorf("Something went wrong")
			}
		}
		if file.GetFileName() == "username" {
			username := file.GetContent()
			if username != "username" {
				t.Errorf("Something went wrong")
			}
		}
	}
	// empty subdirs should be deleted
	_, errEmptyDir := os.Stat(filepath.Join(basePath, storeName+"/extraContent/Subdir/youtube/personal"))
	if errEmptyDir != nil {
		t.Errorf("dir should not exist anymore")
	}
	teardown()
}

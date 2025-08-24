package api

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/Carlo451/vb-password-base-package/passwordstoreFilesystem"

	"time"
)

// createConfigDirWithFiles creates the config directory with given params
func createConfigDirWithFiles(parentDir *passwordstoreFilesystem.PasswordStoreDir, owner, encryptionId string) passwordstoreFilesystem.PasswordStoreDir {
	configDir := parentDir.AddEmpytContentDir("configs")
	ownerFile := passwordstoreFilesystem.NewPasswordstoreContentFile(owner, "owner", *configDir)
	encryptionIdFile := passwordstoreFilesystem.NewPasswordstoreContentFile(encryptionId, "enryptionId", *configDir)
	lastEditedFile := passwordstoreFilesystem.NewPasswordstoreContentFile(time.Now().String(), "lastEdited", *configDir)
	configDir.AppendFiles(ownerFile, encryptionIdFile, lastEditedFile)
	return *parentDir
}

// createRootDir creates a new root directory (like a new vault) a base directory with all relevant configs
func createRootDir(path, name, owner, encryptionId string) passwordstoreFilesystem.PasswordStoreDir {
	rootDir := passwordstoreFilesystem.CreateNewEmptyStoreDir(name, path)
	createConfigDirWithFiles(&rootDir, owner, encryptionId)
	rootDir.WriteDirectory()
	return rootDir
}

// createCustomRootDir creates a new root directory (like a new vault) a base directory with custom configs
func createCustomRootDir(path, name string, configs []passwordstoreFilesystem.PasswordStoreContentFile) passwordstoreFilesystem.PasswordStoreDir {
	rootDir := passwordstoreFilesystem.CreateNewEmptyStoreDir(name, path)
	contentDir := rootDir.AddEmpytContentDir("configs")
	for _, config := range configs {
		contentDir.AppendFile(&config)
	}
	rootDir.WriteDirectory()
	return rootDir
}

// checkIfSubDirPathExistsAndReturnLastSubDir
// runs through the directories and checks if all subdirectories already exists
// returns the last existing subdirectory on the path
// returns if all subdirectories exist
// returns the remaining subdirectories, which are not existing yet
func checkIfSubDirPathExistsAndReturnLastSubDir(dir passwordstoreFilesystem.PasswordStoreDir, orderedSubDirList []string) (passwordstoreFilesystem.PasswordStoreDir, bool, []string) {
	if len(orderedSubDirList) == 0 {
		return dir, true, orderedSubDirList
	}
	if len(dir.GetStoreDirectories()) == 0 {
		return dir, false, orderedSubDirList
	}
	for _, directory := range dir.GetStoreDirectories() {
		if directory.GetDirName() == orderedSubDirList[0] {
			dirList := orderedSubDirList[1:]
			return checkIfSubDirPathExistsAndReturnLastSubDir(directory, dirList)
		}
	}
	return dir, false, orderedSubDirList
}

// createSubDirectoriesRek runs recursively down the given sub directories and creates each one
// dir - the directory in which the sub directory should be created
// subdirs - a ordered list of sub dirs which are given down recursively only the first ohne of the list is taken
// returns the last created subDir
func createSubDirectoriesRek(dir passwordstoreFilesystem.PasswordStoreDir, subdirs []string) passwordstoreFilesystem.PasswordStoreDir {
	if len(subdirs) == 0 {
		return dir
	}
	var newSubDir = passwordstoreFilesystem.CreateNewEmptyStoreDir(subdirs[0], dir.GetAbsoluteDirectoryPath())
	dir.AddDirectory(&newSubDir)
	dir.WriteDirectory()
	dirList := subdirs[1:]
	return createSubDirectoriesRek(newSubDir, dirList)
}

// addAndWriteContentToContentDirectory - appends one new file with the given identifier(fileName) and the content and later saves it to the filesystem
func addAndWriteContentToContentDirectory(content, identifier string, contentDir passwordstoreFilesystem.PasswordStoreContentDir) {
	file := passwordstoreFilesystem.NewPasswordstoreContentFile(content, identifier, contentDir)
	contentDir.AppendFile(file)
	contentDir.WriteDirectory()
}

// writeOrOverwriteFileInContentDir - takes the lastSubDir searches for the correct content dir and overwrites the contentfile or writes the content file
func writeOrOverwriteFileInContentDir(dir passwordstoreFilesystem.PasswordStoreDir, contentDirName, content, identifier string) (bool, error) {
	contentDirs := dir.GetContentDirectories()
	for _, contentDir := range contentDirs {
		if contentDir.GetDirName() == contentDirName {
			contentDir.CreateAndAppend(content, identifier)
			return true, nil
		}
	}
	return false, errors.New("content directory not found")
}

// removeEmptyDirsRecUpWards removes empty directories the tree up
func removeEmptyDirsRecUpWards(lastSubDir passwordstoreFilesystem.PasswordStoreDir) {
	if len(lastSubDir.GetAllDirs()) == 0 {
		passwordstoreFilesystem.RemoveDirectory(&lastSubDir)
		overlayingStore := passwordstoreFilesystem.ReadDirDownFromPath(lastSubDir.GetDirEntryPath())
		removeEmptyDirsRecUpWards(overlayingStore)
	} else {
		return
	}
}

func (h *PasswordStoreHandler) combineRelativePathWithBasePath(relativePath ...string) string {
	removedBasePathOfArgs := []string{h.path}

	// removes the base part if this exists (everything to the base directory)
	for _, arg := range relativePath {
		if strings.Contains(arg, h.path) {
			_, editedArg, _ := strings.Cut(arg, h.path)
			removedBasePathOfArgs = append(removedBasePathOfArgs, editedArg)

		} else {
			removedBasePathOfArgs = append(removedBasePathOfArgs, arg)
		}
	}

	lastString := removedBasePathOfArgs[len(removedBasePathOfArgs)-1]
	for i := len(removedBasePathOfArgs) - 2; i >= 0; i-- {
		if strings.Contains(lastString, removedBasePathOfArgs[i]) {
			_, editedArg, _ := strings.Cut(lastString, removedBasePathOfArgs[i])
			removedBasePathOfArgs[i+1] = editedArg
			lastString = removedBasePathOfArgs[i]
		}
	}

	return filepath.Join(removedBasePathOfArgs...)
}

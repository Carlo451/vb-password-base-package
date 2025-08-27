package api

import (
	"errors"

	"github.com/Carlo451/vb-password-base-package/logger"
	"github.com/Carlo451/vb-password-base-package/passwordstoreFilesystem"
	"github.com/Carlo451/vb-password-base-package/pathparser"
)

type PasswordStoreHandler struct {
	path string
}

func (h *PasswordStoreHandler) GetPath() string {
	return h.path
}

func (h *PasswordStoreHandler) CreatePasswordStore(name, user, encryptionId string) passwordstoreFilesystem.PasswordStoreDir {
	return createRootDir(h.path, name, user, encryptionId)
}

func (h *PasswordStoreHandler) CreateCustomPasswordStore(name string, configs []passwordstoreFilesystem.PasswordStoreContentFile) passwordstoreFilesystem.PasswordStoreDir {
	return createCustomRootDir(h.path, name, configs)
}

func (h *PasswordStoreHandler) ReadPasswordStore(name string) passwordstoreFilesystem.PasswordStoreDir {
	return passwordstoreFilesystem.ReadDirDownFromPath(h.combineRelativePathWithBasePath(name))
}

// AddContentDirectoryToStore Adds a new Content directory with the given content and identifiers as well as the corresponding subdirectories in the path
// path to the last sub dir, where the new contentDir should be created
// storeName - the name of the root password store
// contentDirName - the name of the content directory
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) AddContentDirectoryToStore(path, storeName, contentDirName, content, identifier string) {
	parsedPath := pathparser.ParsePathWithContentDirectory(h.combineRelativePathWithBasePath(storeName), h.combineRelativePathWithBasePath(storeName, path, contentDirName))
	store := h.ReadPasswordStore(storeName)
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
// storeName - the name of the root password store
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) InsertContentInContentDirectory(path, storeName, content, identifier string) (bool, error) {
	dirExists, err := CheckIfContentDirectoryExists(h.combineRelativePathWithBasePath(storeName, path))
	if err != nil {
		return false, err
	}
	if dirExists {
		if CheckIfContentFileExists(h.combineRelativePathWithBasePath(path, identifier)) {
			return false, errors.New("content file already exists, it needs to be updated not inserted")
		}
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(h.combineRelativePathWithBasePath(storeName), h.combineRelativePathWithBasePath(storeName, path))
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			return writeOrOverwriteFileInContentDir(dir, parsedPath.ContentDirectory, content, identifier)
		}
	}
	return false, errors.New("the content directory in the Path does not exist")
}

// UpdateContentInContentDirectory updates the content of a content file - or creates a content file if the one with that identifier does not exist
// path - path to the content Directory
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) UpdateContentInContentDirectory(path, storeName, content, identifier string) (bool, error) {
	dirExists, err := CheckIfContentDirectoryExists(h.combineRelativePathWithBasePath(storeName, path))
	if err != nil {
		logger.ApiLogger.Error("the content directory does not exist")
		return false, err
	}
	if dirExists {
		if !CheckIfContentFileExists(h.combineRelativePathWithBasePath(path, identifier)) {
			logger.ApiLogger.Warn("the file doesn't exist yet but it will be created")
			return h.InsertContentInContentDirectory(path, storeName, content, identifier)
		}
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(h.combineRelativePathWithBasePath(storeName), h.combineRelativePathWithBasePath(storeName, path))
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			return writeOrOverwriteFileInContentDir(dir, parsedPath.ContentDirectory, content, identifier)
		}
	}
	return false, errors.New("the content directory in the Path does not exist")
}

// DeleteContentInContentDirectory deletes content of content dir -> if that is the only input in the dir, all underlying dir which are empty are removed too
// path to the content directory
// storeName
// identifier name of the file
func (h *PasswordStoreHandler) DeleteContentInContentDirectory(path, storeName, identifier string) error {
	if !CheckIfContentFileExists(h.combineRelativePathWithBasePath(storeName, path, identifier)) {
		logger.ApiLogger.Warn("the file doesn't exist")
		return error(errors.New("the file doesn't exist"))
	}
	contentDir, _ := h.ReadContentDir(h.combineRelativePathWithBasePath(path), storeName)
	contentDir.RemoveFileWithFilename(identifier)
	if len(contentDir.ReturnFiles()) == 0 {
		deleted, _ := h.RemoveDirectory(path, storeName, true)
		if !deleted {
			return error(errors.New("the now empty ContentDirectory could not be deleted"))
		}
	}
	return nil
}

// RemoveDirectory removes a directory
// path to the directory
// storeName
// removeSubDirs whether the directories up the tree should be deleted too if they are empty
func (h *PasswordStoreHandler) RemoveDirectory(path, storeName string, removeSubDirs bool) (bool, error) {
	dirExists := CheckIfDirectoryExists(h.combineRelativePathWithBasePath(storeName, path))
	if dirExists {
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(h.combineRelativePathWithBasePath(storeName), h.combineRelativePathWithBasePath(storeName, path))
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			for _, contentDir := range dir.GetAllDirs() {
				if contentDir.GetDirName() == parsedPath.ContentDirectory {
					passwordstoreFilesystem.RemoveDirectory(contentDir)
					if removeSubDirs {
						updatedDir := passwordstoreFilesystem.ReadDirDownFromPath(dir.GetAbsoluteDirectoryPath())
						removeEmptyDirsRecUpWards(updatedDir)
					}
					return true, nil
				}
			}
		}
	}
	logger.ApiLogger.Error("the directory does not exist")
	return false, nil
}

// MoveDirectory moves a content directory to another sub directory
// path is the path to the current content directory
// storeName is the name of the store
// pathToNewSubDirectory is the path of the sub directory where the content dir should be moved
// it gets the content of the directory and writes them to another directory
func (h *PasswordStoreHandler) MoveDirectory(path, storeName, pathToNewSubDirectory string) (bool, error) {
	contentDir, err := h.ReadContentDir(path, storeName)
	if err != nil {
		return false, err
	}
	for _, file := range contentDir.ReturnFiles() {
		h.AddContentDirectoryToStore(pathToNewSubDirectory, storeName, contentDir.GetDirName(), file.GetContent(), file.GetFileName())
	}
	h.RemoveDirectory(path, storeName, true)
	return true, nil
}

// ReadContentDir reads and returns the a content directory
// path to the content directory
// storeName
func (h *PasswordStoreHandler) ReadContentDir(path, storeName string) (*passwordstoreFilesystem.PasswordStoreContentDir, error) {
	dirExists, err := CheckIfContentDirectoryExists(h.combineRelativePathWithBasePath(storeName, path))
	if err != nil {
		return nil, err
	}
	if dirExists {
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(h.combineRelativePathWithBasePath(storeName), h.combineRelativePathWithBasePath(storeName, path))
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			for _, contentDir := range dir.GetContentDirectories() {
				if contentDir.GetDirName() == parsedPath.ContentDirectory {
					return &contentDir, nil
				}
			}
		}
	}
	return nil, errors.New("the content directory in the Path does not exist")
}

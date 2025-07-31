package api

import (
	"errors"
	"github.com/Carlo451/vb-password-base-package/logger"
	"github.com/Carlo451/vb-password-base-package/passwordstore/passwordstoreFilesystem"
	"github.com/Carlo451/vb-password-base-package/pathparser"
	"path/filepath"
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
	return passwordstoreFilesystem.ReadDirDownFromPath(filepath.Join(h.path, name))
}

// AddContentDirectoryToStore Adds a new Content directory with the given content and identifiers as well as the corresponding subdirectories in the path
// path - path to the last sub dir, where the new contentDir should be created
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// contentDirName - the name of the content directory
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) AddContentDirectoryToStore(path, storeName, contentDirName, content, identifier string) {
	parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(h.path, storeName), filepath.Join(path, contentDirName))
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
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) InsertContentInContentDirectory(path, storeName, content, identifier string) (bool, error) {
	dirExists, err := CheckIfContentDirectoryExists(path)
	if err != nil {
		return false, err
	}
	if dirExists {
		if CheckIfContentFileExists(filepath.Join(path, identifier)) {
			return false, errors.New("content file already exists, it needs to be updated not inserted")
		}
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(h.path, storeName), path)
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			return writeOrOvewriteFileInContentDir(dir, parsedPath.ContentDirectory, content, identifier)
		}
	}
	return false, errors.New("the content directory in the Path does not exist")
}

// UpdateContentInContentDirectory updates the content of a content file - or creates a content file if the one with that identifer does not exist
// path - path to the content Directory
// basePath - path to the Password Store Base - the directory where all Stores are stored
// storeName - the name of the root password store
// content - the content
// identifier - the name of the file
func (h *PasswordStoreHandler) UpdateContentInContentDirectory(path, storeName, content, identifier string) (bool, error) {
	dirExists, err := CheckIfContentDirectoryExists(path)
	if err != nil {
		logger.ApiLogger.Error("the content directory does not exist")
		return false, err
	}
	if dirExists {
		if !CheckIfContentFileExists(filepath.Join(path, identifier)) {
			logger.ApiLogger.Warn("the file doesn't exist yet but it will be created")
			return h.InsertContentInContentDirectory(path, storeName, content, identifier)
		}
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(h.path, storeName), path)
		dir, exists, _ := checkIfSubDirPathExistsAndReturnLastSubDir(store, parsedPath.SubDirectories)
		if exists {
			return writeOrOvewriteFileInContentDir(dir, parsedPath.ContentDirectory, content, identifier)
		}
	}
	return false, errors.New("the content directory in the Path does not exist")
}

func (h *PasswordStoreHandler) RemoveDirectory(path, storeName string, removeSubDirs bool) (bool, error) {
	dirExists := CheckIfDirectoryExists(path)
	if dirExists {
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(h.path, storeName), path)
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

func (h *PasswordStoreHandler) ReadContentDir(path, storeName string) (*passwordstoreFilesystem.PasswordStoreContentDir, error) {
	dirExists, err := CheckIfContentDirectoryExists(path)
	if err != nil {
		return nil, err
	}
	if dirExists {
		store := h.ReadPasswordStore(storeName)
		parsedPath := pathparser.ParsePathWithContentDirectory(filepath.Join(h.path, storeName), path)
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

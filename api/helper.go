package api

import (
	"errors"
	"vb-password-store-base/passwordstore/passwordstoreFilesystem"
)

func CreateRootDir(path, name, owner, encryptionId string) passwordstoreFilesystem.PasswordStoreDir {
	rootDir := passwordstoreFilesystem.CreateNewEmptyStoreDir(name, path)
	rootDir.CreateConfigDirWithFiles(owner, encryptionId)
	rootDir.WriteDirectory()
	return rootDir
}

func CreateNotLoadedRootDir(path, name string) *passwordstoreFilesystem.PasswordStoreDir {
	rootDir := passwordstoreFilesystem.CreateNewEmptyStoreDir(name, path)
	return &rootDir
}

// runs through the directories and checks if all subdirectories already exists
// returns the last existing subdirectory on the path
// returns if all subdirectories exist
// returns the remaning subdirectories, which are not existing yet
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

// runs recursive through a list of subdir strings and creates plus write this into the filesystem
// returns the last created subDir
func createSubDirectoriesRek(dir passwordstoreFilesystem.PasswordStoreDir, subdirs []string) passwordstoreFilesystem.PasswordStoreDir {
	if len(subdirs) == 0 {
		return dir
	}
	var newSubDir = passwordstoreFilesystem.CreateNewEmptyStoreDir(subdirs[0], dir.GetAbsoluteDirectoryPath())
	dir.AddSubDirectory(newSubDir)
	dir.WriteDirectory()
	dirList := subdirs[1:]
	return createSubDirectoriesRek(newSubDir, dirList)
}

func addAndWriteContentToContentDirectory(content, identifier string, contentDir passwordstoreFilesystem.PasswordStoreContentDir) {
	file := passwordstoreFilesystem.NewPasswordstoreContentFile(content, identifier, contentDir)
	contentDir.AppendFile(file)
	contentDir.WriteDirectory()
}

func writeOrOverrideFileInContentDir(dir passwordstoreFilesystem.PasswordStoreDir, contentDirName, content, identifier string) (bool, error) {
	contentDirs := dir.GetContentDirectories()
	for _, contentDir := range contentDirs {
		if contentDir.GetDirName() == contentDirName {
			contentDir.CreateAndAppend(content, identifier)
			return true, nil
		}
	}
	return false, errors.New("content directory not found")
}

/*func removeEmptyDirsRecUpWards(lastSubDir passwordstoreFilesystem.PasswordStoreDir) {
	if len(lastSubDir.GetAllDirs()) < 2 {
		passwordstoreFilesystem.RemoveDirectory(&lastSubDir)
		removeEmptyDirsRecUpWards(lastSubDir.Ge)
	} else {
		return
	}
}*/

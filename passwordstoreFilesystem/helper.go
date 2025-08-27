package passwordstoreFilesystem

import (
	"os"
	"path/filepath"
)

func (dir *PasswordStoreDir) checkIfUnderlayingDirIsContentDir(entry os.DirEntry) bool {
	newDirShouldBeContentDir := false
	entrys, error := os.ReadDir(filepath.Join(dir.GetAbsoluteDirectoryPath(), entry.Name()))
	if error != nil {
		panic(error)
	}
	if len(entrys) == 0 {
		return newDirShouldBeContentDir
	}
	if !entrys[0].IsDir() {
		newDirShouldBeContentDir = true
	}
	return newDirShouldBeContentDir
}

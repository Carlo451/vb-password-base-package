package io

import (
	"os"
	"path/filepath"
	"vb-password-store-base/interfaces"
)

func WriteFileContents(file interfaces.PasswordStoreFile) {
	os.WriteFile(filepath.Join(file.GetUnderlyingDirectoryPath(), file.GetFileName()), []byte(file.ReturnFileContents()), 0644)
}

func WriteDirectory(dir interfaces.PasswordStoreDir) {
	os.Mkdir(filepath.Join(dir.GetUnderlyingDirectoryPath(), dir.GetDirName()), 0775)
}

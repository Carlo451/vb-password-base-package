package io

import (
	"os"
	"path/filepath"
	"vb-password-store-base/interfaces"
)

func WriteFileContents(file interfaces.PasswordStoreFile) {
	os.WriteFile(filepath.Join(file.GetDirectoryPath(), file.GetName()), []byte(file.ReturnFileContents()), 0644)
}

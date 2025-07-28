package passwordstoreFilesystem

import (
	"os"
)

func WriteFileContents(file File) {
	os.WriteFile(file.GetAbsolutePath(), []byte(file.GetContent()), 0644)
}

func WriteDirectory(dir Directory) {
	os.Mkdir(dir.GetAbsoluteDirectoryPath(), 0775)
}

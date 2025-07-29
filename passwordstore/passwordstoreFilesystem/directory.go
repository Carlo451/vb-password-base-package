package passwordstoreFilesystem

import (
	"os"
)

type Directory interface {
	GetDirName() string
	GetAbsoluteDirectoryPath() string
	ReadDirectoryRec(entry os.DirEntry)
	ReadDirectory() Directory
	WriteDirectory()
}

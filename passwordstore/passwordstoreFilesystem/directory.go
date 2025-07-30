package passwordstoreFilesystem

import (
	"os"
)

type Directory interface {
	GetDirName() string
	GetAbsoluteDirectoryPath() string
	GetDirEntryPath() string
	ReadDirectoryRec(entry os.DirEntry)
	ReadDirectory() Directory
	WriteDirectory()
}

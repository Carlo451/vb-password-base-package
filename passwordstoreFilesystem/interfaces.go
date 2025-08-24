package passwordstoreFilesystem

import (
	"os"
)

type Directory interface {
	// GetDirName - returns the name of the directory
	GetDirName() string

	// GetAbsoluteDirectoryPath returns the complete absolut path of that directory
	GetAbsoluteDirectoryPath() string

	// GetDirEntryPath - returns the path of the directory in which this particular directory is saved
	GetDirEntryPath() string

	// ReadDirectoryRec takes a DireEntry and is used when ReadDirectory is called Recursively
	ReadDirectoryRec(entry os.DirEntry)

	// ReadDirectory reads the calling directory and every directory down the tree
	// returns ther calling directory with content
	ReadDirectory() Directory

	// WriteDirectory - writes the calling directory into the filesystem
	WriteDirectory()
}

type File interface {
	// GetFileName - returns the filename
	GetFileName() string

	// GetDirectoryName - returns the directory name of the directory it is saved in
	GetDirectoryName() string

	// GetDirectory - returns the directory it is saved in
	GetDirectory() Directory

	// GetAbsolutePath - returns the full absolut path of the file
	GetAbsolutePath() string

	// WriteContent - writes the content of the file into the filesystem
	WriteContent()

	// SetContent - sets the content of the file
	SetContent(string)

	// GetContent - returns the content of the file
	GetContent() string

	// SetDirectory - sets the directory of the file
	SetDirectory(Directory)
}

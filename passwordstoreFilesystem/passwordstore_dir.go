package passwordstoreFilesystem

import (
	"os"
	"path/filepath"
)

type PasswordStoreDir struct {
	directoryName string
	directoryPath string
	directories   []Directory
}

// GetDirEntryPath - returns the path in which this particular directory is saved
func (dir *PasswordStoreDir) GetDirEntryPath() string {
	return dir.directoryPath
}

// GetDirName - returns the name of the directory
func (p PasswordStoreDir) GetDirName() string {
	return p.directoryName
}

// GetAbsoluteDirectoryPath returns the complete absolut path of that directory
func (p PasswordStoreDir) GetAbsoluteDirectoryPath() string {
	return filepath.Join(p.directoryPath, p.directoryName)
}

// ReadDirectoryRec takes a DireEntry and is used when ReadDirectory is called Recursively
func (p *PasswordStoreDir) ReadDirectoryRec(entry os.DirEntry) {
	if entry.IsDir() {
		if p.checkIfUnderlayingDirIsContentDir(entry) {
			p.directories = append(p.directories, NewCleanContentDirectory(p, entry).ReadDirectory())
		} else {
			p.directories = append(p.directories, NewCleanPasswordStoreDir(p, entry).ReadDirectory())
		}
	}
}

// ReadDirectory reads that particular directory and everything what is down the tree
func (p *PasswordStoreDir) ReadDirectory() Directory {
	ReadDir(p)
	return p
}

// WriteDirectory writes that particular directory and everything what is down the tree to the filesystem
func (p *PasswordStoreDir) WriteDirectory() {
	WriteDirectory(p)
	for _, directory := range p.directories {
		directory.WriteDirectory()
	}
}

// GetStoreDirectories returns all directories which are no content directories
func (p *PasswordStoreDir) GetStoreDirectories() []PasswordStoreDir {
	dirs := []PasswordStoreDir{}
	for _, directory := range p.directories {
		ok, err := directory.(*PasswordStoreDir)
		if err == true {
			dirs = append(dirs, *ok)
		}
	}
	return dirs
}

// GetContentDirectories returns all directories which are content directories
func (p *PasswordStoreDir) GetContentDirectories() []PasswordStoreContentDir {
	var dirs []PasswordStoreContentDir
	for _, directory := range p.directories {
		ok, err := directory.(*PasswordStoreContentDir)
		if err == true {
			dirs = append(dirs, *ok)
		}
	}
	return dirs
}

// GetAllDirs returns all directories of that particular directory
func (p *PasswordStoreDir) GetAllDirs() []Directory {
	return p.directories
}

// AddDirectory adds a new directory to the directory list, but it will not be written automatically
func (p *PasswordStoreDir) AddDirectory(directory Directory) {
	p.directories = append(p.directories, directory)
}

// AddEmpytContentDir adds a new empty content directory to the directory
func (p *PasswordStoreDir) AddEmpytContentDir(name string) *PasswordStoreContentDir {
	contentDir := NewEmptyContentDirecotry(*p, name)
	p.AddDirectory(&contentDir)
	return &contentDir
}

// CreateNewEmptyStoreDir creates an empty store directory with the path of the directory where the new one will be stored
func CreateNewEmptyStoreDir(name, path string) PasswordStoreDir {
	return PasswordStoreDir{
		name, path, []Directory{},
	}
}

// NewCleanPasswordStoreDir creates a object of the directory entry
func NewCleanPasswordStoreDir(headDir *PasswordStoreDir, entry os.DirEntry) *PasswordStoreDir {
	return &PasswordStoreDir{
		directoryName: entry.Name(),
		directoryPath: headDir.GetAbsoluteDirectoryPath(),
	}
}

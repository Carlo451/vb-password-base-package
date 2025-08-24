package passwordstoreFilesystem

import (
	"errors"
	"os"
	"path/filepath"
)

type PasswordStoreContentDir struct {
	directoryName string
	directoryPath string
	contents      []File
}

// GetDirEntryPath - returns the path in which this particular directory is saved
func (p *PasswordStoreContentDir) GetDirEntryPath() string {
	return p.directoryPath
}

// GetDirName - returns the name of the directory
func (p *PasswordStoreContentDir) GetDirName() string {
	return p.directoryName
}

// GetAbsoluteDirectoryPath - returns the
func (p *PasswordStoreContentDir) GetAbsoluteDirectoryPath() string {
	return filepath.Join(p.directoryPath, p.directoryName)
}

// ReadDirectoryRec - since this is normally called by the filesystem reader and it is the leaf of a tree it calls WriteDirectory
func (p *PasswordStoreContentDir) ReadDirectoryRec(entry os.DirEntry) {
	p.ReadDirectory()
}

// ReadDirectory - reads the content out of the filesystem for every file
func (p *PasswordStoreContentDir) ReadDirectory() Directory {
	p.contents = ReadFiles(*p)
	return p
}

// WriteDirectory writes the directory and all the files of that directory into the filesystem
func (p *PasswordStoreContentDir) WriteDirectory() {
	WriteDirectory(p)
	p.WriteFiles()
}

// WriteFile writes a single file into the filesystem
func (p *PasswordStoreContentDir) WriteFile(file File) {
	WriteFileContents(file)
}

// WriteFiles writes all files into the filysystem
func (p *PasswordStoreContentDir) WriteFiles() {
	for _, file := range p.contents {
		p.WriteFile(file)
	}
}

// AppendFile appends a new file to the list and sets it directory to the calling one
func (p *PasswordStoreContentDir) AppendFile(file File) {
	file.SetDirectory(p)
	p.contents = append(p.contents, file)
}

// AppendFiles append multiple files
func (p *PasswordStoreContentDir) AppendFiles(files ...File) {
	for _, file := range files {
		p.AppendFile(file)
	}
}

// CreateAndAppend creates a new file in the directory and sets its content
func (p *PasswordStoreContentDir) CreateAndAppend(content, identifier string) {
	var file File = NewPasswordstoreContentFile(content, identifier, *p)
	p.AppendFile(file)
	p.WriteFiles()
}

// RemoveFile removes a file from the directory
func (p *PasswordStoreContentDir) RemoveFile(f File) {
	for i, file := range p.contents {
		if file.GetFileName() == f.GetFileName() {
			p.contents = append(p.contents[:i], p.contents[i+1:]...)
		}
	}
}

// RemoveFileWithFilename removes a file from the directory with filename
func (p *PasswordStoreContentDir) RemoveFileWithFilename(fileName string) {
	for i, file := range p.contents {
		if file.GetFileName() == fileName {
			p.contents = append(p.contents[:i], p.contents[i+1:]...)
		}
	}
}

// ReturnFiles returns the files of the directory
func (p *PasswordStoreContentDir) ReturnFiles() []File {
	return p.contents
}

// ReturnFile returns the file with the specific name if exists
func (p *PasswordStoreContentDir) ReturnFile(fileName string) (File, error) {
	for _, file := range p.contents {
		if file.GetFileName() == fileName {
			return file, nil
		}
	}
	return nil, errors.New("file not found")
}

// LookUpFile checks if file with that name exists
func (p *PasswordStoreContentDir) LookUpFile(fileName string) bool {
	for _, file := range p.ReturnFiles() {
		if file.GetFileName() == fileName {
			return true
		}
	}
	return false
}

// NewCleanContentDirectory creates a new ContentDirectory with the parent dir and the corresponding filesystem entry
func NewCleanContentDirectory(dir Directory, entry os.DirEntry) *PasswordStoreContentDir {
	return &PasswordStoreContentDir{
		directoryName: entry.Name(),
		directoryPath: dir.GetAbsoluteDirectoryPath(),
	}
}

// NewEmptyContentDirecotry creates a new content directory with the parent directory and the name of the content dir
func NewEmptyContentDirecotry(dir PasswordStoreDir, name string) PasswordStoreContentDir {
	return PasswordStoreContentDir{
		directoryName: name,
		directoryPath: dir.GetAbsoluteDirectoryPath(),
		contents:      []File{},
	}
}

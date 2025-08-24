package passwordstoreFilesystem

import "path/filepath"

type PasswordStoreContentFile struct {
	content   string
	name      string
	directory Directory
}

// SetDirectory - sets the directory of the file
func (p *PasswordStoreContentFile) SetDirectory(directory Directory) {
	p.directory = directory
}

// GetFileName - returns the filename
func (p *PasswordStoreContentFile) GetFileName() string {
	return p.name
}

// GetDirectoryName -returns dir name
func (p *PasswordStoreContentFile) GetDirectoryName() string {
	return p.directory.GetDirName()
}

// GetDirectory - returns the directory it is saved in
func (p *PasswordStoreContentFile) GetDirectory() Directory {
	return p.directory
}

// GetAbsolutePath - returns the full absolut path of the file
func (p *PasswordStoreContentFile) GetAbsolutePath() string {
	return filepath.Join(p.directory.GetAbsoluteDirectoryPath(), p.GetFileName())
}

// WriteContent - writes the content of the file into the filesystem
func (p *PasswordStoreContentFile) WriteContent() {
	WriteFileContents(p)
}

// SetContent - sets the content of the file
func (p *PasswordStoreContentFile) SetContent(s string) {
	p.content = s
}

// GetContent - returns the content of the file
func (p *PasswordStoreContentFile) GetContent() string {
	return p.content
}

// NewPasswordstoreContentFile creates a new content File
func NewPasswordstoreContentFile(content string, name string, directory PasswordStoreContentDir) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		content:   content,
		name:      name,
		directory: &directory,
	}
}

// NewCleanPasswordStoreContentFile creates a new clean content file -> contains no data yet
func NewCleanPasswordStoreContentFile(name string, contentDir PasswordStoreContentDir) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		directory: &contentDir,
		name:      name,
	}
}

// NewCleanUnrelatedPasswordStoreContentFile creates a new clean and unrelated file obj, that means that file has no relation to any directory or data
func NewCleanUnrelatedPasswordStoreContentFile(name string) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		name: name,
	}
}

// NewUnrelatedPasswordStoreContentFile creates a new content file obj with content, but it is unrelated to any directory
func NewUnrelatedPasswordStoreContentFile(content, name string) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		name:    name,
		content: content,
	}
}

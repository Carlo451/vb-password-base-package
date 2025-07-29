package passwordstoreFilesystem

import "path/filepath"

type File interface {
	GetFileName() string
	GetDirectoryName() string
	GetDirectory() Directory
	GetAbsolutePath() string
	WriteContent()
	SetContent(string)
	GetContent() string
}

type PasswordStoreContentFile struct {
	content   string
	name      string
	directory Directory
}

func (p *PasswordStoreContentFile) GetFileName() string {
	return p.name
}

func (p *PasswordStoreContentFile) GetDirectoryName() string {
	return p.directory.GetDirName()
}

func (p *PasswordStoreContentFile) GetDirectory() Directory {
	return p.directory
}

func (p *PasswordStoreContentFile) GetAbsolutePath() string {
	return filepath.Join(p.directory.GetAbsoluteDirectoryPath(), p.GetFileName())
}

func (p *PasswordStoreContentFile) WriteContent() {
	WriteFileContents(p)
}

func (p *PasswordStoreContentFile) SetContent(s string) {
	p.content = s
}

func (p *PasswordStoreContentFile) GetContent() string {
	return p.content
}

func NewPasswordstoreContentFile(content string, name string, directory PasswordStoreContentDir) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		content:   content,
		name:      name,
		directory: &directory,
	}
}

func NewCleanPasswordStoreContentFile(name string, contentDir PasswordStoreContentDir) *PasswordStoreContentFile {
	return &PasswordStoreContentFile{
		directory: &contentDir,
		name:      name,
	}
}

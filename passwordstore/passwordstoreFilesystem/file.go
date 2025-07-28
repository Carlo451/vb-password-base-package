package passwordstoreFilesystem

import "path/filepath"

type File interface {
	GetFileName() string
	GetDirectoryName() string
	GetDirectory() Directory
	GetAbsolutePath() string

	WriteContent() string
	SetContent(string)
	GetContent() string
	NewFile() File
}

type PasswordStoreContentFile struct {
	content   string
	name      string
	directory Directory
}

func (p PasswordStoreContentFile) GetFileName() string {
	return p.name
}

func (p PasswordStoreContentFile) GetDirectoryName() string {
	return p.directory.GetDirName()
}

func (p PasswordStoreContentFile) GetDirectory() Directory {
	return p.directory
}

func (p PasswordStoreContentFile) GetAbsolutePath() string {
	return filepath.Join(p.GetAbsolutePath(), p.GetFileName())
}

func (p PasswordStoreContentFile) WriteContent() string {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreContentFile) SetContent(s string) {
	p.content = s
}

func (p PasswordStoreContentFile) GetContent() string {
	return p.content
}

func (p PasswordStoreContentFile) NewFile() File {
	//TODO implement me
	panic("implement me")
}

func NewPasswordStoreContentFile(content string, name string, directory Directory) PasswordStoreContentFile {
	return PasswordStoreContentFile{
		content:   content,
		name:      name,
		directory: directory,
	}
}

func NewCleanPasswordStoreContentFile(name string, contentDir PasswordStoreContentDir) PasswordStoreContentFile {
	return PasswordStoreContentFile{
		directory: contentDir,
		name:      name,
	}
}

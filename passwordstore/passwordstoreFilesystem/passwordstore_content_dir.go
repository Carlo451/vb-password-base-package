package passwordstoreFilesystem

import (
	"os"
	"path/filepath"
)

type PasswordStoreContentDir struct {
	directoryName string
	directoryPath string
	contents      []File
}

func (p *PasswordStoreContentDir) GetDirName() string {
	return p.directoryName
}

func (p *PasswordStoreContentDir) GetAbsoluteDirectoryPath() string {
	return filepath.Join(p.directoryPath, p.directoryName)
}

func (p *PasswordStoreContentDir) ReadDirectoryRec(entry os.DirEntry) {
	p.ReadDirectory()
	test := entry.Name()
	if test == "." {
	}
}

func (p *PasswordStoreContentDir) ReadDirectory() Directory {
	p.contents = ReadFiles(*p)
	return p
}

func (p *PasswordStoreContentDir) WriteDirectory() {
	WriteDirectory(p)
	p.WriteFiles()
}

func (p *PasswordStoreContentDir) WriteFile(file File) {
	WriteFileContents(file)
}

func (p *PasswordStoreContentDir) WriteFiles() {
	for _, file := range p.contents {
		p.WriteFile(file)
	}
}

func (p *PasswordStoreContentDir) AppendFile(file File) {
	p.contents = append(p.contents, file)
}

func (p *PasswordStoreContentDir) AppendFiles(files ...File) {
	for _, file := range files {
		p.AppendFile(file)
	}
}

func (p *PasswordStoreContentDir) CreateAndAppend(content, identifier string) {
	var file File = NewPasswordstoreContentFile(content, identifier, *p)
	p.AppendFile(file)
	p.WriteFiles()
}

func (p *PasswordStoreContentDir) RemoveFile(f File) {
	for i, file := range p.contents {
		if file.GetFileName() == f.GetFileName() {
			p.contents = append(p.contents[:i], p.contents[i+1:]...)
		}
	}
}

func NewCleanContentDirectory(dir Directory, entry os.DirEntry) *PasswordStoreContentDir {
	return &PasswordStoreContentDir{
		directoryName: entry.Name(),
		directoryPath: dir.GetAbsoluteDirectoryPath(),
	}
}

func NewEmptyContentDirecotry(dir PasswordStoreDir, name string) PasswordStoreContentDir {
	return PasswordStoreContentDir{
		directoryName: name,
		directoryPath: dir.GetAbsoluteDirectoryPath(),
		contents:      []File{},
	}
}

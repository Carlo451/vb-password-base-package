package passwordstoreFilesystem

import (
	"os"
	"path/filepath"
)

type Directory interface {
	GetDirName() string
	GetAbsoluteDirectoryPath() string
	ReadDirectoryRec(entry os.DirEntry)
	ReadDirectory() Directory
	AddDirectoryToFileSystem()
	WriteDirectory()
}

type PasswordStoreDir struct {
	directoryName string
	directoryPath string
	directories   []Directory
}

func NewCleanPasswordStoreDir(headDir PasswordStoreDir, entry os.DirEntry) *PasswordStoreDir {
	return &PasswordStoreDir{
		directoryName: entry.Name(),
		directoryPath: filepath.Join(headDir.GetAbsoluteDirectoryPath(), entry.Name()),
	}
}

func (p PasswordStoreDir) GetDirName() string {
	return p.GetDirName()
}

func (p PasswordStoreDir) GetAbsoluteDirectoryPath() string {
	return p.GetAbsoluteDirectoryPath()
}

func (p PasswordStoreDir) ReadDirectoryRec(entry os.DirEntry) {
	if entry.IsDir() {
		if p.checkIfUnderlayingDirIsContentDir(entry) {
			p.directories = append(p.directories, NewCleanContentDirectory(p, entry).ReadDirectory())
		} else {
			p.directories = append(p.directories, NewCleanPasswordStoreDir(p, entry).ReadDirectory())
		}
	}
}

func (p PasswordStoreDir) ReadDirectory() Directory {
	ReadDir(p)
	return p
}

func (p PasswordStoreDir) AddDirectoryToFileSystem() {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreDir) WriteDirectory() {
	//TODO implement me
	panic("implement me")
}

type PasswordStoreContentDir struct {
	directoryName string
	directoryPath string
	contents      []File
}

func (p PasswordStoreContentDir) GetDirName() string {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreContentDir) GetAbsoluteDirectoryPath() string {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreContentDir) ReadDirectoryRec(entry os.DirEntry) {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreContentDir) ReadDirectory() Directory {
	p.contents = ReadFiles(p)
	return p
}

func (p PasswordStoreContentDir) AddDirectoryToFileSystem() {
	//TODO implement me
	panic("implement me")
}

func (p PasswordStoreContentDir) WriteDirectory() {
	//TODO implement me
	panic("implement me")
}

func NewCleanContentDirectory(dir Directory, entry os.DirEntry) PasswordStoreContentDir {
	return PasswordStoreContentDir{
		directoryName: entry.Name(),
		directoryPath: filepath.Join(dir.GetAbsoluteDirectoryPath(), entry.Name()),
	}
}

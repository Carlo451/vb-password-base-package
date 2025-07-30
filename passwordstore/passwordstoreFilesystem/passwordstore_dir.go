package passwordstoreFilesystem

import (
	"os"
	"path/filepath"
	"time"
)

type PasswordStoreDir struct {
	directoryName string
	directoryPath string
	directories   []Directory
}

func (dir *PasswordStoreDir) GetDirEntryPath() string {
	return dir.directoryPath
}

func NewCleanPasswordStoreDir(headDir *PasswordStoreDir, entry os.DirEntry) *PasswordStoreDir {
	return &PasswordStoreDir{
		directoryName: entry.Name(),
		directoryPath: headDir.GetAbsoluteDirectoryPath(),
	}
}

func (p PasswordStoreDir) GetDirName() string {
	return p.directoryName
}

func (p PasswordStoreDir) GetAbsoluteDirectoryPath() string {
	return filepath.Join(p.directoryPath, p.directoryName)
}

func (p *PasswordStoreDir) ReadDirectoryRec(entry os.DirEntry) {
	if entry.IsDir() {
		if p.checkIfUnderlayingDirIsContentDir(entry) {
			p.directories = append(p.directories, NewCleanContentDirectory(p, entry).ReadDirectory())
		} else {
			p.directories = append(p.directories, NewCleanPasswordStoreDir(p, entry).ReadDirectory())
		}
	}
}

func (p *PasswordStoreDir) ReadDirectory() Directory {
	ReadDir(p)
	return p
}

func (p *PasswordStoreDir) WriteDirectory() {
	WriteDirectory(p)
	for _, directory := range p.directories {
		directory.WriteDirectory()
	}
}

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

func (p *PasswordStoreDir) GetContentDirectories() []PasswordStoreContentDir {
	dirs := []PasswordStoreContentDir{}
	for _, directory := range p.directories {
		ok, err := directory.(*PasswordStoreContentDir)
		if err == true {
			dirs = append(dirs, *ok)
		}
	}
	return dirs
}

func (p *PasswordStoreDir) GetAllDirs() []Directory {
	return p.directories
}

func (p *PasswordStoreDir) AddDirectory(directory Directory) {
	p.directories = append(p.directories, directory)
}

func (p *PasswordStoreDir) AddSubDirectory(directory PasswordStoreDir) {
	p.AddDirectory(&directory)
}

func (p *PasswordStoreDir) AddContentDirectory(directory PasswordStoreContentDir) {
	p.AddDirectory(&directory)
}

/*func (p *PasswordStoreDir) GetUnderlayingSubDirectory() *Directory {
	return ReadDirFromPath(p.directoryPath)
}*/

/*func (p *PasswordStoreDir) FindContentDirectoryByName(name string) PasswordStoreContentDir {
	contentDirs := p.GetContentDirectories()
	for _, dir := range contentDirs {
		if
	}
	for _, directory := range p.directories {

	}
}*/

func (p *PasswordStoreDir) CreateConfigDirWithFiles(owner, encryptionId string) {
	configDir := NewEmptyContentDirecotry(*p, "configs")
	ownerFile := NewPasswordstoreContentFile(owner, "owner", configDir)
	encryptionIdFile := NewPasswordstoreContentFile(encryptionId, "enryptionId", configDir)
	lastEditedFile := NewPasswordstoreContentFile(time.Now().String(), "lastEdited", configDir)
	configDir.AppendFiles(ownerFile, encryptionIdFile, lastEditedFile)
	p.directories = append(p.directories, &configDir)
}

func CreateNewEmptyStoreDir(name, path string) PasswordStoreDir {
	return PasswordStoreDir{
		name, path, []Directory{},
	}
}

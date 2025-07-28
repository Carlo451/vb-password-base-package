package passwordstoreConfig

import (
	"os"
	"path/filepath"
	"vb-password-store-base/passwordstore"
	"vb-password-store-base/passwordstore/passwordstoreFilesystem"
)

type PasswordStoreConfigDir struct {
	name         string
	path         string
	owner        passwordstore.PasswordStoreFile
	readers      passwordstore.PasswordStoreFile
	writers      passwordstore.PasswordStoreFile
	encryptionId passwordstore.PasswordStoreFile
	lastEdited   passwordstore.PasswordStoreFile
}

func (store *PasswordStoreConfigDir) AddDirToFileSystem() {
	passwordstoreFilesystem.WriteDirectory(store)
	passwordstoreFilesystem.WriteFileContents(store.owner)
	passwordstoreFilesystem.WriteFileContents(store.readers)
	passwordstoreFilesystem.WriteFileContents(store.writers)
	passwordstoreFilesystem.WriteFileContents(store.encryptionId)
	passwordstoreFilesystem.WriteFileContents(store.lastEdited)
}

func NewPasswordStoreConfig(path, owner, encryptionId string, readers, writers []string) *PasswordStoreConfigDir {
	var rekPath string = filepath.Join(path, "config")
	var ownerFileObj passwordstore.PasswordStoreFile = NewOwner(rekPath, owner)
	var encryptionFileObj passwordstore.PasswordStoreFile = NewEncryptionId(rekPath, encryptionId)
	var readersFileObj passwordstore.PasswordStoreFile = NewReader(rekPath, readers)
	var writerFileObj passwordstore.PasswordStoreFile = NewWriter(rekPath, writers)
	var lastEditedFileObj passwordstore.PasswordStoreFile = NewLastEdited(rekPath)

	return &PasswordStoreConfigDir{path: path, name: "config", owner: ownerFileObj, encryptionId: encryptionFileObj, readers: readersFileObj, writers: writerFileObj, lastEdited: lastEditedFileObj}
}

func (store *PasswordStoreConfigDir) GetDirName() string {
	return store.name
}

func (store *PasswordStoreConfigDir) GetUnderlyingDirectoryPath() string {
	return store.path
}

func (cfg *PasswordStoreConfigDir) ReadDirectory(entry os.DirEntry) passwordstore.PasswordStoreDir {
	if !entry.IsDir() {
		cfg.owner = passwordstoreFilesystem.ReadFile(cfg.owner)
		cfg.encryptionId = passwordstoreFilesystem.ReadFile(cfg.encryptionId)
		cfg.readers = passwordstoreFilesystem.ReadFile(cfg.readers)
		cfg.writers = passwordstoreFilesystem.ReadFile(cfg.writers)
		cfg.lastEdited = passwordstoreFilesystem.ReadFile(cfg.lastEdited)
	}
	return cfg
}

package passwordstoreComponents

import (
	"path/filepath"
	"vb-password-store-base/passwordstore/io"
	"vb-password-store-base/passwordstore/passwordstoreConfig"
)

type PasswordStoreConfigDir struct {
	name         string
	path         string
	owner        passwordstoreConfig.Owner
	readers      passwordstoreConfig.Reader
	writers      passwordstoreConfig.Writer
	encryptionId passwordstoreConfig.EncryptionId
	lastEdited   passwordstoreConfig.LastEdited
}

func NewPasswordStoreConfig(path, name, owner, encryptionId string, readers, writers []string) *PasswordStoreConfigDir {
	var rekPath string = filepath.Join(path, name)
	var ownerFileObj passwordstoreConfig.Owner = *passwordstoreConfig.NewOwner(rekPath, owner)
	var encryptionFileObj passwordstoreConfig.EncryptionId = *passwordstoreConfig.NewEncryptionId(rekPath, encryptionId)
	var readersFileObj passwordstoreConfig.Reader = *passwordstoreConfig.NewReader(rekPath, readers)
	var writerFileObj passwordstoreConfig.Writer = *passwordstoreConfig.NewWriter(rekPath, writers)
	var lastEditedFileObj passwordstoreConfig.LastEdited = *passwordstoreConfig.NewLastEdited(rekPath)

	return &PasswordStoreConfigDir{name: "config", owner: ownerFileObj, encryptionId: encryptionFileObj, readers: readersFileObj, writers: writerFileObj, lastEdited: lastEditedFileObj}
}

func (cfg *PasswordStoreConfigDir) GetName() string {
	return cfg.name
}

func (cfg *PasswordStoreConfigDir) WriteDirectoryConfigContent() {
	io.WriteFileContents(&cfg.owner)
}

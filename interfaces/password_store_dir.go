package interfaces

type PasswordStoreDir interface {
	GetName() string
	ReturnDirContents() map[string]string
}

type PasswordStoreFile interface {
	GetName() string
	ReturnFileContents() string
	UpdateFileContents(newContent string)
	GetDirectoryPath() string
}

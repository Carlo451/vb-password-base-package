package passwordstoreComponents

type PasswordStoreSubDir struct {
	name        string
	contentDirs []PasswordStoreContentDir
}

func (store *PasswordStoreSubDir) GetName() string {
	return store.name
}

func (cfg *PasswordStoreSubDir) ReturnDirContents() map[string]string {
}

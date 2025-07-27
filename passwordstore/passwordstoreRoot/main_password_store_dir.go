package passwordstoreComponents

type PasswordStoreRootDir struct {
	name    string
	configs PasswordStoreConfigDir
	subDirs []PasswordStoreSubDir
}

func (store *PasswordStoreRootDir) GetName() string {
	return store.name
}

func (cfg *PasswordStoreRootDir) ReturnDirContents() map[string]string {
	return nil
}

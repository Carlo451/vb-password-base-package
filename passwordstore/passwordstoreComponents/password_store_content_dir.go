package passwordstoreComponents

type PasswordStoreContentDir struct {
	name     string
	contents []PasswordStoreContent
}

func (store *PasswordStoreContentDir) GetName() string {
	return store.name
}

func (cfg *PasswordStoreContentDir) ReturnDirContents() map[string]string {
}

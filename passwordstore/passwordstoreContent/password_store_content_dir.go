package passwordstoreComponents

type PasswordStoreContentDir struct {
	name     string
	path     string
	contents []PasswordStoreContent
}

func NewPasswordStoreContentDir(path, name string, contents []PasswordStoreContent) *PasswordStoreContentDir {
	return &PasswordStoreContentDir{name: name, contents: contents, path: path}
}

func (store *PasswordStoreContentDir) GetUnderlyingDirectoryPath() string {
	return store.path
}

func (store *PasswordStoreContentDir) GetDirName() string {
	return store.name
}

func (store *PasswordStoreContentDir) SetContents(content PasswordStoreContent) {
	store.contents = append(store.contents, content)
}

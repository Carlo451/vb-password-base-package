package passwordstoreComponents

import "vb-password-store-base/passwordstore/passwordstoreContent"

type PasswordStoreSubDir struct {
	name        string
	contentDirs []passwordstoreContent.PasswordStoreContentDir
}

func (store *PasswordStoreSubDir) GetName() string {
	return store.name
}

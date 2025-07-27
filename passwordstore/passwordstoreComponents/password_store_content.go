package passwordstoreComponents

type PasswordStoreContent struct {
	path             string
	name             string
	encryptedContent string
}

func (content *PasswordStoreContent) GetName() string {
	return content.name
}

func (content *PasswordStoreContent) ReturnFileContents() string {
	return content.encryptedContent
}

func (content *PasswordStoreContent) GetEncryptedContent() string {
	return content.encryptedContent
}

func (content *PasswordStoreContent) UpdateFileContents(newEncryptedString string) {
	content.encryptedContent = newEncryptedString
}

func (content *PasswordStoreContent) GetDirectoryPath() string {
	return content.path
}

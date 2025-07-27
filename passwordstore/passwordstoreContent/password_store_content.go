package passwordstoreComponents

type PasswordStoreContent struct {
	path             string
	name             string
	encryptedContent string
}

func NewPasswordStoreContent(path string, name string, encryptedContent string) *PasswordStoreContent {
	return &PasswordStoreContent{
		name:             name,
		path:             path,
		encryptedContent: encryptedContent,
	}
}

func CreatePasswordContent(path string, encryptedContent string) *PasswordStoreContent {
	return &PasswordStoreContent{
		name:             "password",
		path:             path,
		encryptedContent: encryptedContent,
	}
}

func (content *PasswordStoreContent) GetFileName() string {
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

func (content *PasswordStoreContent) GetUnderlyingDirectoryPath() string {
	return content.path
}

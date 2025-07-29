package passwordstoreFilesystem

var naming = map[string]string{
	"password": "password",
	"email":    "email",
}

func (p *PasswordStoreContentDir) CreatePasswordFile(encryptedPassword string) PasswordStoreContentFile {
	return PasswordStoreContentFile{
		content:   encryptedPassword,
		name:      naming["password"],
		directory: p,
	}
}

func (p *PasswordStoreContentDir) CreateEmailFile(encryptedEmail string) PasswordStoreContentFile {
	return PasswordStoreContentFile{
		content:   encryptedEmail,
		name:      naming["password"],
		directory: p,
	}
}

/*func CreateAdditionalPasswordFile(number int, additionalPassword string) PasswordStoreContentFile {

}

func CreateUsernamePasswordField(encryptedUsername string) PasswordStoreContentFile {

}

func CreateUrlField(encryptedUrl string) PasswordStoreContentFile {

}

func CreateAdditionalUrlField(encryptedUrl string) PasswordStoreContentFile {

}*/

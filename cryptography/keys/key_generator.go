package keys

import (
	"fmt"
	"log"

	"filippo.io/age"
)

// GenerateAsymmetricKey Generates a KeyPair
func GenerateAsymmetricKey() AsymmetricKeyPair {
	privateKey, err := age.GenerateX25519Identity()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Recipient()
	return NewAsymmetricKeyPair(publicKey.String(), privateKey.String())
}

// GenerateNewKeyPairFromPrivateKey generate a new public key from private key (perhaps to compare public keys and check if private key fits t)
func GenerateNewKeyPairFromPrivateKey(privateKey string) (*AsymmetricKeyPair, error) {
	key, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		return nil, err
	}
	pubKey := key.Recipient().String()
	keyPair := NewAsymmetricKeyPair(pubKey, privateKey)
	return &keyPair, nil
}

// GenerateSymmetricKey Generates a simple standard long key
func GenerateSymmetricKey() (*SymmetricKey, error) {
	return GenerateSymmetricKeyWithSpecialLength(30)
}

// GenerateSymmetricKeyWithSpecialLength Generates a simple key where the user can specify the length of the key
func GenerateSymmetricKeyWithSpecialLength(length int) (*SymmetricKey, error) {
	if length < 8 {
		return nil, fmt.Errorf("the password length is too small, the password must be at least 8 characters")
	}
	var password, _ = generatesRandomKeyString(length)
	return NewSymmetricKey(password), nil
}

// GenerateFormattedStandardKey Generates a simple standard long formatted key with standard size char groups
func GenerateFormattedStandardKey() (*SymmetricKey, error) {
	return GenerateFormattedSymmetricKeyWithLength(20, 5)
}

// GenerateFormattedSymmetricKeyWithLength Generates a key with a specific length and specific length of format char groups. Mind that in order to work the group length must be a divider of the length
func GenerateFormattedSymmetricKeyWithLength(length, groupLength int) (*SymmetricKey, error) {
	if (length%groupLength != 0) || length < groupLength*2 || length%groupLength != 0 {
		return nil, fmt.Errorf("the password length must be at least double the group length long. And can be split up in char groups of %b", groupLength)
	}
	var password, _ = generatesRandomKeyString(length)

	for !checkPasswordForAllChars(password) {
		password, _ = generatesRandomKeyString(length)
	}
	var newFormattedPassword string = ""
	for i := 0; i < len(password); i++ {
		if i == (len(password) - 1) {
			newFormattedPassword += string(password[i])
			break
		}
		if i%groupLength == 0 && i != 0 {
			newFormattedPassword += "-"
			newFormattedPassword += string(password[i])
			continue
		}
		newFormattedPassword += string(password[i])
	}
	return NewSymmetricKey(newFormattedPassword), nil
}

package symmetricCrypto

import (
	"testing"
	"vb-password-store-base/cryptography/cryptographyoperations"
	"vb-password-store-base/cryptography/keygenerator"
)

func TestSymmetricEnAndDecryption(t *testing.T) {
	key, err := keygenerator.GenerateSymmetricKey()
	if err != nil {
		t.Errorf("SymmetricEnAndDecryption failed with error %s", err)
	}
	var message string = "Password123"
	encryptedMessage, _ := cryptographyoperations.EncryptStringSymmetric(message, key.Key)

	decryptedMessage, _ := cryptographyoperations.DecryptStringSymmetric(encryptedMessage, key.Key)
	if decryptedMessage != message {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

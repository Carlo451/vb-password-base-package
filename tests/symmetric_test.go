package tests

import (
	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keys"

	"testing"
)

func TestSymmetricEnAndDecryption(t *testing.T) {
	key, err := keys.GenerateSymmetricKey()
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

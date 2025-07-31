package symmetricCrypto

import (
	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keygenerator"
	"testing"
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

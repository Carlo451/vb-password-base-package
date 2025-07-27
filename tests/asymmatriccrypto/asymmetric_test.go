package asymmatriccrypto

import (
	"testing"
	"vb-password-store-base/cryptography/cryptographyoperations"
	"vb-password-store-base/cryptography/keygenerator"
	"vb-password-store-base/cryptography/keys"
)

func TestAsymmetricEnAndDecryption(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keygenerator.GenerateAsymmetricKey()

	var message string = "Password123"
	encryptedMessage, _ := cryptographyoperations.EncryptStringAsymmetric(message, keyPair.PublicKey)

	decryptedMessage, _ := cryptographyoperations.DecryptStringAsymmetric(encryptedMessage, keyPair.PrivateKey)
	if decryptedMessage != message {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

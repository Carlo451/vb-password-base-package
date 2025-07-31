package asymmatriccrypto

import (
	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keygenerator"
	"github.com/Carlo451/vb-password-base-package/cryptography/keys"
	"testing"
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

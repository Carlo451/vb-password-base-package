package tests

import (
	"testing"

	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyoperations"
	"github.com/Carlo451/vb-password-base-package/cryptography/keys"
)

func TestAsymmetricEnAndDecryption(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keys.GenerateAsymmetricKey()

	var message string = "Password123"
	encryptedMessage, _ := cryptographyoperations.EncryptStringAsymmetric(message, keyPair.PublicKey)

	decryptedMessage, _ := cryptographyoperations.DecryptStringAsymmetric(encryptedMessage, keyPair.PrivateKey)
	if decryptedMessage != message {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

func TestGenerationOfNewPublicKeyFromPrivateKey(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keys.GenerateAsymmetricKey()
	newKeyPair, _ := keys.GenerateNewKeyPairFromPrivateKey(keyPair.PrivateKey)
	if newKeyPair.PublicKey != keyPair.PublicKey {
		t.Errorf("GenerateNewKeyPairFromPrivateKey Failed")
	}
}

func TestValidationOfAsymmetricKeyPairs(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keys.GenerateAsymmetricKey()
	valid, _ := keyPair.CheckIfKeyPairIsValid()
	if !valid {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

func TestValidationOfAsymmetricKeyPairs_WrongPublicKey(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keys.GenerateAsymmetricKey()
	var editedKeyPair keys.AsymmetricKeyPair = keys.NewAsymmetricKeyPair("publicKey", keyPair.PrivateKey)
	valid, err := editedKeyPair.CheckIfKeyPairIsValid()
	if err != nil {
		t.Errorf("Should throw no error, since priv key is right")
	}
	if valid {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

func TestValidationOfAsymmetricKeyPairs_MalformedPrivateKey(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keys.GenerateAsymmetricKey()
	var editedKeyPair keys.AsymmetricKeyPair = keys.NewAsymmetricKeyPair(keyPair.PublicKey, "privateKEy")
	valid, err := editedKeyPair.CheckIfKeyPairIsValid()
	if err == nil {
		t.Errorf("Should throw error, since priv key malformed")
	}
	if valid {
		t.Errorf("AsymmetricEnAndDecryption Failed")
	}
}

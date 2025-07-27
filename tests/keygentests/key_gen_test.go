package keygentests

import (
	"fmt"
	"testing"
	"vb-password-store-base/cryptography/keygenerator"
	"vb-password-store-base/cryptography/keys"
)

func TestGenerateKeyPair(t *testing.T) {
	var keyPair keys.AsymmetricKeyPair = keygenerator.GenerateAsymmetricKey()
	if keyPair.PublicKey == "" {
		t.Errorf("public key not generated")
	}
}

func TestGenerateSymmetricKeyStandard(t *testing.T) {
	keyPoint, err := keygenerator.GenerateSymmetricKey()
	if err != nil {
		t.Error(err)
	}
	var key keys.SymmetricKey = *keyPoint
	if keyPoint == nil {
		t.Errorf("key not generated")
	}
	fmt.Println(key.Key)
}

func TestGenerateSymmetricKeyFormatted(t *testing.T) {
	keyPoint, err := keygenerator.GenerateFormattedSymmetricKeyWithLength(20, 5)
	if err != nil {
		t.Error(err)
	}
	if keyPoint == nil {
		t.Errorf("key not generated")
	}
}

func TestGenerateSymmetricKeyFormattedWrongGroupSize(t *testing.T) {
	keyPoint, err := keygenerator.GenerateFormattedSymmetricKeyWithLength(20, 3)
	if err == nil {
		t.Error(err)
	}
	if keyPoint != nil {
		t.Errorf("key not generated")
	}

}

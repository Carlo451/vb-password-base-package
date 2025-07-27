package cryptographyoperations

import (
	"bytes"
	"io"
	"strings"

	"filippo.io/age"
)

func DecryptStringAsymmetric(encryptedString, privateKey string) (string, error) {
	identity, err := age.ParseX25519Identity(strings.TrimSpace(privateKey))
	if err != nil {
		return "", err
	}

	r := strings.NewReader(encryptedString)

	decryptor, err := age.Decrypt(r, identity)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if _, err := io.Copy(&out, decryptor); err != nil {
		return "", err
	}
	return out.String(), nil
}

func DecryptStringSymmetric(encryptedString, passphrase string) (string, error) {
	identity, err := age.NewScryptIdentity(passphrase)
	if err != nil {
		return "", err
	}
	r := strings.NewReader(encryptedString)

	decryptor, err := age.Decrypt(r, identity)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	_, err = io.Copy(&out, decryptor)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

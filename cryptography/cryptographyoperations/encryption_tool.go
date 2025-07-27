package cryptographyoperations

import (
	"bytes"
	"io"

	"filippo.io/age"
)

func EncryptStringAsymmetric(plaintext, publicKey string) (string, error) {
	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	encryptor, err := age.Encrypt(&out, recipient)
	if err != nil {
		return "", err
	}

	if _, err := io.WriteString(encryptor, plaintext); err != nil {
		return "", err
	}
	if err := encryptor.Close(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func EncryptStringSymmetric(plaintext, passphrase string) (string, error) {
	recipient, err := age.NewScryptRecipient(passphrase)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	encryptor, err := age.Encrypt(&out, recipient)
	if err != nil {
		return "", err
	}

	if _, err := io.WriteString(encryptor, plaintext); err != nil {
		return "", err
	}
	if err := encryptor.Close(); err != nil {
		return "", err
	}

	return out.String(), nil
}

package keys

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/Carlo451/vb-password-base-package/cryptography/cryptographyconsts"
)

// generates a random key string with specific length
func generatesRandomKeyString(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		c, err := secureRandomCharFrom(cryptographyconsts.SecureRandomCharset)
		if err != nil {
			return "", err
		}
		result[i] = c[0]
	}
	return string(result), nil
}

// takes random string of a string
func secureRandomCharFrom(set string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
	if err != nil {
		return "", err
	}
	return string(set[n.Int64()]), nil
}

// checks if string contains minimum one char of each string group
func checkPasswordForAllChars(password string) bool {
	var lowerLettersList []string = strings.Split(cryptographyconsts.LowerLetters, "")
	var upperLettersList []string = strings.Split(cryptographyconsts.UpperLetters, "")
	var digitsList []string = strings.Split(cryptographyconsts.Digits, "")
	var specialCharsList []string = strings.Split(cryptographyconsts.SpecialChars, "")
	if checkPasswordForSpecificChars(password, lowerLettersList) && checkPasswordForSpecificChars(password, upperLettersList) && checkPasswordForSpecificChars(password, digitsList) && checkPasswordForSpecificChars(password, specialCharsList) {
		return true
	}
	return false
}

// check if string contains minimum one char of a specif string group
func checkPasswordForSpecificChars(password string, chars []string) bool {
	for _, char := range chars {
		if strings.Contains(password, char) {
			return true
		}
	}
	return false
}

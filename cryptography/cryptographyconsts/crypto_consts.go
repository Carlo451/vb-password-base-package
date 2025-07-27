package cryptographyconsts

const (
	LowerLetters        = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits              = "0123456789"
	SpecialChars        = "!@#$%^&*()_=+[]{}<>?/|"
	SecureRandomCharset = LowerLetters +
		UpperLetters +
		Digits +
		SpecialChars +
		LowerLetters +
		UpperLetters +
		Digits
)

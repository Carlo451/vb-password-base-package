package keys

type AsymmetricKeyPair struct {
	PublicKey  string
	PrivateKey string
}

func NewAsymmetricKeyPair(pub, priv string) AsymmetricKeyPair {
	return AsymmetricKeyPair{
		PublicKey:  pub,
		PrivateKey: priv,
	}
}

// CheckIfKeyPairIsValid checks if the public key and private key is a valid keyPair
func (key AsymmetricKeyPair) CheckIfKeyPairIsValid() (bool, error) {
	newKeyPair, err := GenerateNewKeyPairFromPrivateKey(key.PrivateKey)
	if err != nil {
		return false, err
	}
	if newKeyPair.PublicKey == key.PublicKey {
		return true, nil
	}
	return false, nil
}

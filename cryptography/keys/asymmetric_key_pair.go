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

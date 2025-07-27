package keys

type SymmetricKey struct {
	Key string
}

func NewSymmetricKey(key string) *SymmetricKey {
	return &SymmetricKey{Key: key}
}

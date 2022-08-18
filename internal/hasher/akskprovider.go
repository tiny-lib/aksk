package hasher

type AkSKHashHelper interface {
	HashWithKey(data []byte, key []byte) []byte
	VerifyHash(data []byte, key []byte, signatureToCompareWith []byte) bool
}

package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
)

type Sha256Hash struct {
}

func (s *Sha256Hash) HashWithKey(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (s *Sha256Hash) VerifyHash(data []byte, key []byte, signatureToCompareWith []byte) bool {
	signature := s.HashWithKey(data, key)
	return subtle.ConstantTimeCompare(signature, signatureToCompareWith) == 1
}

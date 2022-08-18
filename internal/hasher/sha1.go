package hasher

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
)

type Sha1Hash struct {
}

func (s *Sha1Hash) HashWithKey(data []byte, key []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func (s *Sha1Hash) VerifyHash(data []byte, key []byte, signatureToCompareWith []byte) bool {
	signature := s.HashWithKey(data, key)
	return subtle.ConstantTimeCompare(signature, signatureToCompareWith) == 1
}

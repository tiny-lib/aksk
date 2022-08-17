package Signer

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Sha256(msg []byte) []byte {
	sh256 := sha256.New()
	sh256.Write(msg)

	return sh256.Sum(nil)
}

func Sha256WithKey(data []byte, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

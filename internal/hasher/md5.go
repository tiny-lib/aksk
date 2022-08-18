package hasher

import (
	"crypto/md5"
)

// Md5Hash data
func Md5Hash(data []byte) []byte {
	mac := md5.New()
	mac.Write(data)
	return mac.Sum(nil)
}

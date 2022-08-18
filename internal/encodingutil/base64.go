package encodingutil

import "encoding/base64"

// Base64Encode encode payload to string
func Base64Encode(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

// Base64Decode decode payload to []byte
func Base64Decode(payload string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(payload)
}

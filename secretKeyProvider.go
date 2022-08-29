package aksk

// SecretKeyProvider provide an interface to get user's secretKey via
// accessKey (such as db or file)
type SecretKeyProvider interface {
	GetSecretKey(accessKey string) (secretKey string, err error)
}

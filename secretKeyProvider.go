package aksk

type SecretKeyProvider interface {
	GetSecretKey(accessKey string) (secretKey string, err error)
}

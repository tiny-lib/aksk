package aksk

import "github.com/czyt/aksk/internal/hasher"

const (
	defaultAuthorization = "X-API-KEY"
	defaultTimeStampKey  = "ts"
)

// Option is ak/sk option.
type Option func(*options)

// options
type options struct {
	baseAuthHeaderKey string
	timeStampKey      string
	encodeUrl         bool
	hashHelper        hasher.AkSKHashHelper
	secretKeyProvider SecretKeyProvider
}

// WithAuthorizationHeader set 	the authorization header
func WithAuthorizationHeader(header string) Option {
	return func(o *options) {
		o.baseAuthHeaderKey = header
	}
}

// WithEncodeUrl set whether to encode url
func WithEncodeUrl(encodeUrl bool) Option {
	return func(o *options) {
		o.encodeUrl = encodeUrl
	}
}

// WithTimeStampKey custom the WithTimeStamp key to fetch the timestamp
func WithTimeStampKey(timeStamp string) Option {
	return func(o *options) {
		o.timeStampKey = timeStamp
	}
}

// WithSecretKeyProvider set the provider to get the secret key
func WithSecretKeyProvider(provider SecretKeyProvider) Option {
	return func(o *options) {
		o.secretKeyProvider = provider
	}
}

// WithHashHelper set the hash helper for hash logic
func WithHashHelper(helper hasher.AkSKHashHelper) Option {
	return func(o *options) {
		o.hashHelper = helper
	}
}

func applyDefaultOptions() Option {
	return func(o *options) {
		if o.baseAuthHeaderKey == "" {
			o.baseAuthHeaderKey = defaultAuthorization
		}
		if o.timeStampKey == "" {
			o.timeStampKey = defaultTimeStampKey
		}
		if o.hashHelper == nil {
			o.hashHelper = &hasher.Sha1Hash{}
		}
	}
}

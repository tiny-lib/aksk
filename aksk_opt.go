package aksk

const (
	defaultAuthorization = "Authorization"
	defaultExpiresKey    = "expires"
	defaultSignatureKey  = "signature"
)

// Option is ak/sk option.
type Option func(*options)

// options
type options struct {
	accessKey     string
	secretKey     string
	authHeaderKey string
	expiresKey    string
	signatureKey  string
	encodeUrl     bool
	extraPayload  map[string]interface{}
}

// WithAccessKey set accessKey
func WithAccessKey(accessKey string) Option {
	return func(o *options) {
		o.accessKey = accessKey
	}
}

// WithSecretKey set secretKey
func WithSecretKey(secretKey string) Option {
	return func(o *options) {
		o.secretKey = secretKey
	}
}

// WithAuthorizationHeader set 	the authorization header
func WithAuthorizationHeader(header string) Option {
	return func(o *options) {
		o.authHeaderKey = header
	}
}

// WithEncodeUrl set whether to encode url
func WithEncodeUrl(encodeUrl bool) Option {
	return func(o *options) {
		o.encodeUrl = encodeUrl
	}
}

// WithExtraPayload set extraPayload to allow user put more info in the auth logic
func WithExtraPayload(extraPayload map[string]interface{}) Option {
	return func(o *options) {
		o.extraPayload = extraPayload
	}
}

// WithExpiresKey custom the expires key
func WithExpiresKey(expiresKey string) Option {
	return func(o *options) {
		o.expiresKey = expiresKey
	}
}

// WithSignatureKey custom the signature Key
func WithSignatureKey(signatureKey string) Option {
	return func(o *options) {
		o.signatureKey = signatureKey
	}
}

func applyDefaultOptions() Option {
	return func(o *options) {
		if o.authHeaderKey == "" {
			o.authHeaderKey = defaultAuthorization
		}
		if o.expiresKey == "" {
			o.expiresKey = defaultExpiresKey
		}
		if o.signatureKey == "" {
			o.signatureKey = defaultSignatureKey
		}
	}
}

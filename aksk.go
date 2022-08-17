package aksk

import (
	"context"
	"github.com/czyt/aksk/internal/builderPool"
	"github.com/czyt/aksk/internal/encoder"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	// reason holds the error reason.
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingAuthorizationHeader = errors.Unauthorized(reason, "Authorization Header is missing")
	ErrSecretKeyProviderNotSet    = errors.Unauthorized(reason, "Secret Key Provider Not Set")
	ErrAuthorizationInvalid       = errors.Unauthorized(reason, "Authorization is invalid")
	ErrAuthorizationExpired       = errors.Unauthorized(reason, "Authorization has expired")
	ErrAuthorizationParseFail     = errors.Unauthorized(reason, "Fail to parse Authorization")
	ErrUnSupportSigningMethod     = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext               = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider          = errors.Unauthorized(reason, "Authorization provider is missing")
	ErrSignToken                  = errors.Unauthorized(reason, "Can not sign Authorization.Is the key correct?")
	ErrGetKey                     = errors.Unauthorized(reason, "Can not get key while signing Authorization")
)

func Server(opts ...Option) middleware.Middleware {
	o := &options{}
	opts = append(opts, applyDefaultOptions())
	for _, opt := range opts {
		opt(o)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if o.secretKeyProvider == nil {
				return nil, ErrSecretKeyProviderNotSet
			}
			if tr, ok := transport.FromServerContext(ctx); ok {
				// check header with authorizationKey
				authorization := tr.RequestHeader().Get(o.authHeaderKey)
				if authorization == "" {
					return nil, ErrMissingAuthorizationHeader
				}
				// Authorization: Access=[Access Key Id], ExpireTime=[expire_time], Signature=[signature]
				// exist:check for signature hex(crypt(Secret Access Key+ExpireTime+URL) ))
				builder := builderPool.New()
				builder.WriteString(o.accessKey)
				if hr, ok := tr.(*http.Transport); ok {
					if o.encodeUrl {
						builder.WriteString(encoder.UrlEncode(hr.Request().RequestURI))
					} else {
						builder.WriteString(hr.Request().RequestURI)
					}

				}
				// not match :return unauthorized

			}
			return handler(ctx, req)
		}

	}
}

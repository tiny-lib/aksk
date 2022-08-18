package aksk

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrMissingAuthorizationHeader = errors.Unauthorized(reason, "Authorization Header is missing")
	ErrSecretKeyProviderNotSet    = errors.Unauthorized(reason, "Secret Key Provider Not Set")
	ErrHashHelperNotSet           = errors.Unauthorized(reason, "Hash Helper Not Set")
	ErrSignCheckFailed            = errors.Unauthorized(reason, "Sign Check failed")
	ErrAuthorizationInvalid       = errors.Unauthorized(reason, "Authorization is invalid")
	ErrAuthorizationExpired       = errors.Unauthorized(reason, "Authorization has expired")
	ErrAuthorizationParseFail     = errors.Unauthorized(reason, "Fail to parse Authorization")
	ErrUnSupportSigningMethod     = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext               = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider          = errors.Unauthorized(reason, "Authorization provider is missing")
	ErrSignToken                  = errors.Unauthorized(reason, "Can not sign Authorization.Is the key correct?")
	ErrGetKey                     = errors.Unauthorized(reason, "Can not get key while signing Authorization")
)

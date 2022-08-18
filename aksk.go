package aksk

import (
	"bytes"
	"context"
	"github.com/czyt/aksk/internal/encodingutil"
	"github.com/czyt/aksk/internal/header"
	"github.com/czyt/aksk/internal/signer"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"io"
)

const (
	// reason holds the error reason.
	reason            string = "UNAUTHORIZED"
	ContentTypeHeader string = "Content-Type"
)

// refer https://www.ietf.org/rfc/rfc2104.txt

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
				var (
					headerKey      = ""
					accessKey      = ""
					secretKey      = ""
					requestUrl     = ""
					httpVerb       = ""
					requestContent = make([]byte, 0)
					contentType    = ""
					unixTimeStamp  = ""
					err            error
				)

				// check header with authorizationKey
				headerKey, accessKey = header.GetAccessKeyFromDynamicHeaderKey(tr.RequestHeader(), o.baseAuthHeaderKey)

				if headerKey == "" {
					return nil, ErrMissingAuthorizationHeader
				}
				originSign := tr.RequestHeader().Get(headerKey)
				if originSign == "" {
					return nil, ErrAuthorizationInvalid
				}
				targetSign, err := encodingutil.Base64Decode(originSign)
				if err != nil {
					return nil, err
				}

				unixTimeStamp = tr.RequestHeader().Get(o.timeStampKey)
				contentType = tr.RequestHeader().Get(ContentTypeHeader)

				if hr, ok := tr.(*http.Transport); ok {
					if o.encodeUrl {
						requestUrl = encodingutil.UrlEncode(hr.Request().RequestURI)
					} else {
						requestUrl = hr.Request().RequestURI
					}
					requestContent, err = io.ReadAll(hr.Request().Body)
					if err != nil {
						return nil, err
					}
					// put the body back
					hr.Request().Body = io.NopCloser(bytes.NewReader(requestContent))
					httpVerb = hr.Request().Method
				}

				// get secret key by access key
				secretKey, err = o.secretKeyProvider.GetSecretKey(accessKey)
				if err != nil {
					return nil, err
				}

				sign := signer.New([]byte(secretKey),
					o.hashHelper,
					signer.WithHttpVerb(httpVerb),
					signer.WithContent(requestContent),
					signer.WithContentType(contentType),
					signer.WithUnixTimeStamp(unixTimeStamp),
					signer.WithRequestUrl(requestUrl),
				)
				checkPass, err := sign.CheckSignValid(targetSign)
				if err != nil {
					return nil, err
				}
				if !checkPass {
					return nil, ErrSignCheckFailed
				}

			}
			return handler(ctx, req)
		}

	}
}

package aksk

import (
	"encoding/hex"
	"github.com/czyt/aksk/internal/builderPool"
	"github.com/czyt/aksk/internal/hasher"
)

type SignGenerator struct {
	requestUrl    string
	httpVerb      string
	contentType   string
	unixTimeStamp string
	secretKey     []byte
	content       []byte

	hashHelper hasher.AkSKHashHelper
}

type Opt func(generator *SignGenerator)

func WithRequestUrl(requestUrl string) Opt {
	return func(generator *SignGenerator) {
		generator.requestUrl = requestUrl
	}
}

func WithHttpVerb(httpVerb string) Opt {
	return func(generator *SignGenerator) {
		generator.httpVerb = httpVerb
	}
}

func WithContent(content []byte) Opt {
	return func(generator *SignGenerator) {
		generator.content = content
	}
}

func WithContentType(contentType string) Opt {
	return func(generator *SignGenerator) {
		generator.contentType = contentType
	}
}

func WithUnixTimeStamp(unixTimeStamp string) Opt {
	return func(generator *SignGenerator) {
		generator.unixTimeStamp = unixTimeStamp
	}
}

func New(secretKey []byte, hashHelper hasher.AkSKHashHelper, options ...Opt) *SignGenerator {
	signer := &SignGenerator{
		secretKey:  secretKey,
		hashHelper: hashHelper,
	}
	for _, option := range options {
		option(signer)
	}
	return signer
}

func (g *SignGenerator) GetSignContent() (string, error) {
	// SignGenerator
	//
	//	 Authorization = AuthorizationHeader + " " + AccessKeyId + ":" + Signature;
	//
	//		Signature = Base64( HashMethod( UTF-8-Encoding-Of(YourAccessKey), UTF-8-Encoding-Of( StringToSign ) ) );
	//
	//		StringToSign = HTTP-Verb + "\n" +
	//		Content-MD5 + "\n" +
	//		Content-Type + "\n" +
	//		UnixTimeStamp + "\n" +
	//		requestURL;
	//
	// /*
	if g.hashHelper == nil {
		return "", ErrHashHelperNotSet
	}
	builder := builderPool.New()
	defer builderPool.Release(builder)
	builder.WriteString(g.httpVerb)
	builder.WriteString("\n")
	hash := hasher.Md5Hash(g.content)
	contentMd5 := hex.EncodeToString(hash)
	builder.WriteString(contentMd5)
	builder.WriteString("\n")
	builder.WriteString(g.contentType)
	builder.WriteString("\n")
	builder.WriteString(g.unixTimeStamp)
	builder.WriteString("\n")
	builder.WriteString(g.requestUrl)
	return builder.String(), nil
}

func (g *SignGenerator) Calculate() ([]byte, error) {

	content, err := g.GetSignContent()
	if err != nil {
		return nil, err
	}
	sign := g.hashHelper.HashWithKey([]byte(content), g.secretKey)
	return sign, nil
}

func (g *SignGenerator) CheckSignValid(targetSign []byte) (bool, error) {
	content, err := g.GetSignContent()
	if err != nil {
		return false, err
	}
	return g.hashHelper.VerifyHash([]byte(content), g.secretKey, targetSign), nil
}

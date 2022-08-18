# aksk
access key/secret key auth middleware for kratos

the logic behind can refer [ amazon s3 authentication](https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3_Authentication2.html)
```javascript

Authorization = AuthorizationHeader + " " + AccessKeyId + ":" + Signature;

Signature = Base64( HashMethod( UTF-8-Encoding-Of(YourSecretKey), UTF-8-Encoding-Of( StringToSign ) ) );

StringToSign = HTTP-Verb + "\n" +
	Content-MD5 + "\n" +
	Content-Type + "\n" +
    UnixTimeStamp + "\n" +
	requestURL;


```
the client `unixTimestamp` should be passed from request header in ms. the middleware has some options:
```javascript
    baseAuthHeaderKey string
	timeStampKey      string
	encodeUrl         bool
	hashHelper        hasher.AkSKHashHelper
	secretKeyProvider SecretKeyProvider
```

default options:
+ hashHelper `Sha1`
+ baseAuthHeader  `X-API-KEY`
+ timestampKey `ts`
+ url not encoded `encodeUrl = false`

reference:
+ https://github.com/dynuc/gophercloud/blob/master/auth/aksk/aksk_signer.go

+ https://blog.csdn.net/makenothing/article/details/81158481

+ https://docs.emqx.com/zh/fabric/latest/api_auth/ak_sk.html#%E4%BD%BF%E7%94%A8ak-sk

+ https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3_Authentication2.html
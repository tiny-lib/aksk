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
the client `unixTimestamp` should be passed from request header in ms.
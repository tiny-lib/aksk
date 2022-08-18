package header

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"regexp"
)

func GetDynamicHeaderKey(header transport.Header, baseHeaderKey string) string {
	headerKeys := header.Keys()
	match, err := regexp.Compile(fmt.Sprintf("^%s.+?", baseHeaderKey))
	if err != nil {
		return ""
	}
	for _, key := range headerKeys {
		if match.MatchString(key) {
			return key
		}
	}
	return ""
}

func GetAccessKeyFromDynamicHeaderKey(header transport.Header, baseHeaderKey string) (key string, value string) {
	match, err := regexp.Compile(fmt.Sprintf("(?m)%s(.+?$)", baseHeaderKey))
	if err != nil {
		return "", ""
	}
	headerKey := GetDynamicHeaderKey(header, baseHeaderKey)
	if match.MatchString(headerKey) {
		groupCapture := match.FindAllStringSubmatch(key, -1)[0]
		if len(groupCapture) == 2 {
			return headerKey, groupCapture[1]
		}
	}
	return headerKey, ""
}

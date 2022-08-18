package header

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexpGroupCapture(t *testing.T) {
	var re = regexp.MustCompile(`(?m)API-KEY(.+?$)`)
	var str = `API-KEY 2121`
	fmt.Println(re.MatchString(str))
	for i, match := range re.FindAllStringSubmatch(str, -1) {
		fmt.Println(i, "len：", len(match))
		for _, s := range match {
			fmt.Println(s)
		}
	}
}

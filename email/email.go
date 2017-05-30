package email

import (
	"regexp"
)

var (
	validFormat = regexp.MustCompile(`[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[\w!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?`)
)

func IsValid(v string) bool {
	return validFormat.MatchString(v)
}

package stringutils

import (
	"regexp"
	"strings"
)

var (
	RegexPhone = regexp.MustCompile(`^1\d{9}\d$`)
	RegexEmail = regexp.MustCompile(`^[\w_-]+(?:\.[\w_-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?$`)
)

func ToString(obj interface{}) string {
	if nil == obj {
		return ""
	}

	var result string
	switch obj.(type) {
	case string:
		result, _ = obj.(string)
	}

	return result
}

func EmptyDefault(obj interface{}, defaultValue string) string {
	value := ToString(obj)
	if strings.TrimSpace(value) == "" {
		return defaultValue
	}

	return value
}

//手机号是否合法
func IsLegalPhoneNumber(phone string) bool {
	return RegexPhone.MatchString(phone)
}

//邮箱是否合法
func IsLegalEmail(phone string) bool {
	return RegexEmail.MatchString(phone)
}

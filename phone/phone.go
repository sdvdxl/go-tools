package phone

import "regexp"

var (
	phoneFormat   map[Country]*regexp.Regexp
	defaultFormat = China
)

const (
	China = Country("China")
)

type Country string

func init() {
	phoneFormat = make(map[Country]*regexp.Regexp)
	phoneFormat[China] = regexp.MustCompile(`^1\d{9}\d$`)
}

func Default(defaultFormat Country) {
	defaultFormat = defaultFormat
}

// IsLegal 手机号是否合法
func IsValid(phone string) bool {
	return phoneFormat[defaultFormat].MatchString(phone)
}

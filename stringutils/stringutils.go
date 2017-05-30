package stringutils

import (
	"strings"
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

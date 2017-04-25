package number

import (
	"fmt"
	"strconv"
)

// DefaultInt 转换对象为 int
func DefaultInt(obj interface{}, defaultValue int) int {
	return int(DefaultInt64(obj, int64(defaultValue)))
}

// DefaultInt64 转换对象为int64
func DefaultInt64(obj interface{}, defaultValue int64) int64 {
	returnValue, ok := obj.(int64)
	if ok {
		return returnValue
	}

	returnValue, err := strconv.ParseInt(fmt.Sprint(obj), 10, 64)
	if err != nil {
		return defaultValue
	}

	return returnValue

}

// DefaultFloat64 转换对象为 float64
func DefaultFloat64(obj interface{}, defaultValue float64) float64 {
	returnValue, ok := obj.(float64)
	if ok {
		return returnValue
	}

	returnValue, err := strconv.ParseFloat(fmt.Sprint(obj), 64)
	if err != nil {
		return defaultValue
	}

	return returnValue
}

// DefaultBool 转换对象为 bool
func DefaultBool(obj interface{}, defaultValue bool) bool {
	returnValue, ok := obj.(bool)
	if ok {
		return returnValue
	}

	returnValue, err := strconv.ParseBool(fmt.Sprint(obj))
	if err != nil {
		return defaultValue
	}

	return returnValue
}

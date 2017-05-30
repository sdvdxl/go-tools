package values

import (
	"reflect"

	"github.com/jinzhu/copier"
)

func IsSameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

// CopyProperties 拷贝属性
// 如果 source 或者 target 有null，那么不进行任何操作
// 否则 把source的属性值拷贝给target
func CopyProperties(source, target interface{}) error {
	if source == nil || target == nil {
		return nil
	}

	return copier.Copy(target, source)

}

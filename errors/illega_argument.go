package errors

// IlleagleArgument 参数不合法
type IlleagleArgument struct {
	CodeError
}

func (e IlleagleArgument) Error() string {
	return e.CodeError.Error()
}

// NewIlleagleArgument 新建
func NewIlleagleArgument(code int, msgs ...interface{}) error {
	return IlleagleArgument{NewCode(code, msgs...)}
}

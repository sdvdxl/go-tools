package errors

// IllegalArgument 参数不合法
type IllegalArgument struct {
	CodeError
}

func (e IllegalArgument) Error() string {
	return e.CodeError.Error()
}

// NewIllegalArgument 新建
func NewIllegalArgument(code int, msgs ...interface{}) error {
	return IllegalArgument{NewCode(code, msgs...)}
}

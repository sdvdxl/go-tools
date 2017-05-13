package errors

// IlleagleArgumentError 参数不合法
type IlleagleArgumentError struct {
	CodeError
}

func (e IlleagleArgumentError) Error() string {
	return e.CodeError.Error()
}

// NewIlleagleArgumentError 新建
func NewIlleagleArgumentError(code int, msgs ...interface{}) error {
	return IlleagleArgumentError{NewCodeError(code, msgs)}
}

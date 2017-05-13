package errors

// NotFoundError 找不到
type NotFoundError struct {
	CodeError
}

func (e NotFoundError) Error() string {
	return e.CodeError.Error()
}

// NewNotFoundError 新建
func NewNotFoundError(code int, msgs ...interface{}) error {
	return NotFoundError{NewCodeError(code, msgs)}
}

package errors

// NotFound 找不到
type NotFound struct {
	CodeError
}

func (e NotFound) Error() string {
	return e.CodeError.Error()
}

// NewNotFound 新建
func NewNotFound(code int, msgs ...interface{}) error {
	return NotFound{NewCode(code, msgs...)}
}

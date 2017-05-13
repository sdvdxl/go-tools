package errors

// NewDuplicatedError 新建
func NewDuplicatedError(code int, msgs ...interface{}) error {
	return DuplicatedError{NewCodeError(code, msgs)}
}

// DuplicatedError 重复
type DuplicatedError struct {
	CodeError
}

func (e DuplicatedError) Error() string {
	return e.CodeError.Error()
}

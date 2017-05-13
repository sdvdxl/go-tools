package errors

// NewDuplicatedError 新建
func NewDuplicatedError(code int, msgs ...interface{}) error {
	return Duplicated{NewCode(code, msgs)}
}

// Duplicated 重复
type Duplicated struct {
	CodeError
}

func (e Duplicated) Error() string {
	return e.CodeError.Error()
}

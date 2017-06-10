package errors

// Unauthorized 权限错误
type Unauthorized struct {
	CodeError
}

func (e Unauthorized) Error() string {
	return e.CodeError.Error()
}

// NewUnauthorized 新建
func NewUnauthorized(code int, msgs ...interface{}) error {
	return Unauthorized{NewCode(code, msgs...)}
}

// Forbidden 权限错误
type Forbidden struct {
	CodeError
}

func (e Forbidden) Error() string {
	return e.CodeError.Error()
}

// NewForbidden 新建
func NewForbidden(code int, msgs ...interface{}) error {
	return Forbidden{NewCode(code, msgs...)}
}

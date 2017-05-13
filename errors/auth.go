package errors

// AuthError 权限错误
type AuthError struct {
	CodeError
}

func (e AuthError) Error() string {
	return e.CodeError.Error()
}

// NewAuthError 新建
func NewAuthError(code int, msgs ...interface{}) error {
	return AuthError{NewCodeError(code, msgs)}
}

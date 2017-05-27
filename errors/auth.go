package errors

// Auth 权限错误
type Auth struct {
	CodeError
}

func (e Auth) Error() string {
	return e.CodeError.Error()
}

// NewAuth 新建
func NewAuth(code int, msgs ...interface{}) error {
	return Auth{NewCode(code, msgs...)}
}

package errors

import (
	"strings"
)

// Unauthorized 权限错误
type Unauthorized struct {
	codeError
}

func (e Unauthorized) Error() string {
	return e.codeError.Error()
}

// NewUnauthorized 新建
func NewUnauthorized(code int, msgs ...string) error {
	return Unauthorized{codeError{code: code, msg: strings.Join(msgs, ",")}}
}

func (e Unauthorized) AddMsg(msgs ...string) Unauthorized {
	e.msg += strings.Join(msgs, ",")
	return e
}

// Forbidden 权限错误
type Forbidden struct {
	codeError
}

func (e Forbidden) Error() string {
	return e.codeError.Error()
}

// NewForbidden 新建
func NewForbidden(code int, msgs ...string) error {
	return Forbidden{codeError{code: code, msg: strings.Join(msgs, ",")}}
}

func (e Forbidden) AddMsg(msgs ...string) Forbidden {
	e.msg += strings.Join(msgs, ",")
	return e
}

package errors

import (
	"strings"
)

// NotFound 找不到
type NotFound struct {
	codeError
}

func (e NotFound) Error() string {
	return e.codeError.Error()
}

// NewNotFound 新建
func NewNotFound(code int, msgs ...string) NotFound {
	return NotFound{codeError{code: code, msg: strings.Join(msgs, ",")}}
}

func (e NotFound) AddMsg(msgs ...string) NotFound {
	e.msg += strings.Join(msgs, ",")
	return e
}

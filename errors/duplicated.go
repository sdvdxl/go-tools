package errors

import (
	"strings"
)

// NewDuplicated 新建
func NewDuplicated(code int, msgs ...string) Duplicated {
	return Duplicated{codeError{code: code, msg: strings.Join(msgs, ",")}}
}

// Duplicated 重复
type Duplicated struct {
	codeError
}

func (e Duplicated) Error() string {
	return e.codeError.Error()
}

func (e Duplicated) AddMsg(msgs ...string) Duplicated {
	e.msg += strings.Join(msgs, ",")
	return e
}

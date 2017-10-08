package errors

import (
	"strings"
)

// IllegalArgument 参数不合法
type IllegalArgument struct {
	codeError
}

func (e IllegalArgument) Error() string {
	return e.codeError.Error()
}

// NewIllegalArgument 新建
func NewIllegalArgument(code int, msgs ...string) IllegalArgument {
	return IllegalArgument{codeError{code: code, msg: strings.Join(msgs, ",")}}
}

func (e IllegalArgument) AddMsg(msgs ...string) IllegalArgument {
	e.msg += strings.Join(msgs, ",")
	return e
}

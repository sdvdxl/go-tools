package errors

import (
	"fmt"
	"strings"
)

// NewCode 新建一个error
func NewCode(code int, msgs ...interface{}) CodeError {
	return CodeError{Code: code, Msg: fmt.Sprintf(strings.Repeat("%s", len(msgs)), msgs...)}
}

// ConstError 不可变error
type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

// CodeError 带错误码的错误
type CodeError struct {
	Log  string `json:"-"` // record log, not show for user
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (e CodeError) Error() string {
	return fmt.Sprint("code:", e.Code, " msg:", e.Msg)
}

var (
	NotFoundError        = NewNotFound(0, "not found")
	DuplicatedError      = NewDuplicatedError(0, "already exists")
	IllegalArgumentError = NewIllegalArgument(0, "illegal argument")
)

// Panic 如果err 不是nil，则panic
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

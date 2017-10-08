package errors

import (
	"fmt"
	"strings"
)

// NewCode 新建一个error
func NewCode(code int, msgs ...string) CodeError {
	return CodeError{Code: code, Msg: strings.Join(msgs, ",")}
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

func (e CodeError) GetLog() string {
	return e.Log
}

func (e CodeError) GetCode() int {
	return e.Code
}

func (e CodeError) GetMsg() string {
	return e.Msg
}

func (e CodeError) Error() string {
	return fmt.Sprint("code:", e.Code, " msg:", e.Msg)
}

func (e CodeError) AddMsg(msgs ...string) CodeError {
	e.Msg += strings.Join(msgs, ",")
	return e
}

type codeError struct {
	log  string
	msg  string
	code int
}

func (e codeError) GetLog() string {
	return e.log
}

func (e codeError) GetCode() int {
	return e.code
}

func (e codeError) GetMsg() string {
	return e.msg
}

func (e codeError) Error() string {
	return fmt.Sprint("code:", e.code, " msg:", e.msg)
}

func (e codeError) AddMsg(msgs ...string) codeError {
	e.msg += strings.Join(msgs, ",")
	return e
}

var (
	NotFoundError        = NewNotFound(0, "not found")
	DuplicatedError      = NewDuplicated(0, "already exists")
	IllegalArgumentError = NewIllegalArgument(0, "illegal argument")
)

// Panic 如果err 不是nil，则panic
// 如果msgs 有填写，则panic msgs+err.Error()
func Panic(err error, msgs ...string) {
	if err != nil {
		if len(msgs) > 0 {
			panic(strings.Join(msgs, ",") + err.Error())
		} else {
			panic(err)
		}
	}
}

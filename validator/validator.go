package validator

import (
	"errors"
	"fmt"

	vt "gopkg.in/validator.v2"
	"uke.cloud/travel/util"
)

var (
	// Validator 校验用
	Validator = vt.NewValidator()
)

func init() {
	Validator.SetValidationFunc("isValidPhone", IsValidPhone)
}

// IsValidPhone 校验是否是合法手机号
func IsValidPhone(v interface{}, param string) error {
	if !util.IsValidPhone(fmt.Sprint(v)) {
		return errors.New("phone is invalid")
	}

	return nil
}

// Valid 根据tags校验
func Valid(v interface{}, tags string) error {
	return vt.Valid(v, tags)
}

package validator

import (
	vt "gopkg.in/validator.v2"
)

var (
	// Validator 校验用
	Validator = vt.NewValidator()
)

// Valid 根据tags校验
func Valid(v interface{}, tags string) error {
	return vt.Valid(v, tags)
}

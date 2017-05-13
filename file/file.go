package file

import (
	"gopkg.in/h2non/filetype.v0"
)

// IsJPG 是不是 jpg
func IsJPG(buf []byte) bool {
	return filetype.IsMIME(buf, "image/jpeg")
}

// IsPNG 是不是png
func IsPNG(buf []byte) bool {
	return filetype.IsMIME(buf, "image/png")
}

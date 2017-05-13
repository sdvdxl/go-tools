package img

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/sdvdxl/go-tools/errors"
	"github.com/sdvdxl/go-tools/file"
)

const (
	// ErrUnsupportedImageType 不支持的图片格式
	ErrUnsupportedImageType = errors.ConstError("unsupported image type")
)

// Crop 剪切图片
func Crop(r io.Reader, x, y, width, height int) (*bytes.Buffer, error) {
	var err error
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	io.Copy(buf, r)

	var rImg image.Image

	if file.IsJPG(buf.Bytes()) {
		rImg, err = jpeg.Decode(buf)
		if err != nil {
			return nil, err
		}
	} else if file.IsPNG(buf.Bytes()) {
		rImg, err = png.Decode(buf)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, ErrUnsupportedImageType
	}

	buf.Truncate(0)
	err = png.Encode(buf, rImg)

	img, _, _ := image.Decode(buf)

	subImg := img.(*image.RGBA64).SubImage(image.Rect(x, y, width+x, height+y))
	err = png.Encode(buf, subImg)

	return buf, err
}

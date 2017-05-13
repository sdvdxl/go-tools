package charset

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

var (
	GB18030   = simplifiedchinese.GB18030
	GBK       = simplifiedchinese.GBK
	HZGB2312  = simplifiedchinese.HZGB2312
	EUCJP     = japanese.EUCJP
	SHIFTJIS  = japanese.ShiftJIS
	ISO2022JP = japanese.ISO2022JP
	EUCKR     = korean.EUCKR
	BIG5      = traditionalchinese.Big5
	ISO8859_2 = charmap.ISO8859_2
	ISO8859_3 = charmap.ISO8859_3
	ISO8859_4 = charmap.ISO8859_4
	ISO8859_5 = charmap.ISO8859_5
	ISO8859_6 = charmap.ISO8859_6
	ISO8859_7 = charmap.ISO8859_7
)

func a() {
}

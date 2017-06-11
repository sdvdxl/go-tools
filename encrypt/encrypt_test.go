package encrypt

import "testing"

func TestAES(t *testing.T) {
	text := []byte("1234567890你好")
	key := []byte("12345678901234567890123456789012345678901234567890")
	AesEncrypt(text, key)
	if string(text) != string(AesDecrypt(AesEncrypt(text, key), key)) {
		t.Fail()
	}
}

package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
)

// MD5 将传入的字符串md5后输出
func MD5(text []byte) string {
	result := md5.Sum(text)
	return hex.EncodeToString(result[:])
}

// SHA512 将参数值 hash，得到 sha512值
func SHA512(text []byte) string {
	result := sha512.Sum512(text)
	return hex.EncodeToString(result[:])
}

func genAesKey(keybyte []byte) []byte {
	key := keybyte
	if len(keybyte) < 32 {
		for i := 0; i < 32-len(keybyte); i++ {
			key = append(key, '0')
		}
	} else {
		key = keybyte[:32]
	}

	return key
}

// HashWithSalt 用给定的字符串 salt 将 text 加密并且  hash512
func HashWithSalt(text, salt []byte) string {
	content := []byte(MD5(AesEncrypt(text, salt)))
	return SHA512(append(content, salt...))
}

// AesEncrypt 加密
func AesEncrypt(origData, key []byte) []byte {
	newKey := genAesKey(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, newKey[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

// AesDecrypt 解密
func AesDecrypt(crypted, key []byte) []byte {
	newKey := genAesKey(key)
	block, err := aes.NewCipher(newKey)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, newKey[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

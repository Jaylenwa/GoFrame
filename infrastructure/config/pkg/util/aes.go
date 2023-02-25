package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// ZeroPadding 填充零
func ZeroPadding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{0}, padding) //剩余用0填充
	return append(cipherText, padText...)

}

// ZeroUnPadding 反填充
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func AESEncrypt(text string, key []byte) (string, error) {
	blockSize := aes.BlockSize //AES的分组大小为16位
	src := []byte(text)
	src = ZeroPadding(src, blockSize) //填充
	out := make([]byte, len(src))
	block, err := aes.NewCipher(key) //用aes创建一个加密器cipher
	if err != nil {
		return "", err
	}
	encrypted := cipher.NewCBCEncrypter(block, key) //CBC分组模式加密
	encrypted.CryptBlocks(out, src)                 //对src进行加密，加密结果放到dst里
	return hex.EncodeToString(out), nil
}

func AESDecrypt(text string, key []byte) (string, error) {
	src, err := hex.DecodeString(text) //转为[]byte
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	block, err := aes.NewCipher(key) //用aes创建一个加密器cipher
	if err != nil {
		return "", err
	}
	decrypted := cipher.NewCBCDecrypter(block, key) //CBC分组模式解密
	decrypted.CryptBlocks(out, src)                 //对src进行解密，解密结果放到dst里
	out = ZeroUnPadding(out)                        //反填充
	return string(out), nil
}

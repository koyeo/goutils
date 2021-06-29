package _crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
)

func EncryptBytes(key, plainText []byte) (cipherText []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	plainBytes := []byte(plainText)
	blockSize := block.BlockSize()
	plainBytes = PKCS7Padding(plainBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherText = make([]byte, len(plainBytes))
	blockMode.CryptBlocks(cipherText, plainBytes)
	return
}

func DecryptBytes(key, cipherText []byte) (plainText []byte, err error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	plainData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainData, cipherText)
	plainText = PKCS7UnPadding(plainData)
	return
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, paddingText...)
}

func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unPadding := int(plainText[length-1])
	return plainText[:(length - unPadding)]
}

func EncryptText(salt, text string) (cipher string, err error) {
	bs, err := EncryptBytes([]byte(salt), []byte(text))
	if err != nil {
		return
	}
	cipher = base64.StdEncoding.EncodeToString(bs)
	return
}

func DecryptText(salt, cipher string) (text string, err error) {
	bs, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return
	}
	bs, err = DecryptBytes([]byte(salt), bs)
	if err != nil {
		return
	}
	text = string(bs)
	return
}

func Md5HMac(salt, text []byte) []byte {
	mac := hmac.New(md5.New, salt)
	mac.Write(text)
	return mac.Sum(nil)
}

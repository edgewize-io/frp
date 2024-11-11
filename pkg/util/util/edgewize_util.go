package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	mathrand "math/rand/v2"
)

func Encrypt(src, key string, iv []byte) []byte {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	//data = padding.PKCS7Padding(data, block.BlockSize())
	data = PKCS7Padding(data, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return out
}

func Decrypt(data []byte, key string, iv []byte) []byte {
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	//plaintext = padding.PKCS5UnPadding(plaintext)
	plaintext = PKCS7UnPadding(plaintext, block.BlockSize())
	return plaintext
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func GenerateAesKey(k string) string {
	m := md5.New()
	m.Write([]byte(k))
	return hex.EncodeToString(m.Sum(nil))
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[mathrand.IntN(len(chars))]
	}

	return string(result)
}

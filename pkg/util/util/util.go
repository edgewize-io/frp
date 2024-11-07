// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	mathrand "math/rand/v2"
	"net"
	"strconv"
	"strings"
	"time"
)

// RandID return a rand string used in frp.
func RandID() (id string, err error) {
	return RandIDWithLen(16)
}

// RandIDWithLen return a rand string with idLen length.
func RandIDWithLen(idLen int) (id string, err error) {
	if idLen <= 0 {
		return "", nil
	}
	b := make([]byte, idLen/2+1)
	_, err = rand.Read(b)
	if err != nil {
		return
	}

	id = fmt.Sprintf("%x", b)
	return id[:idLen], nil
}

func GetAuthKey(token string, timestamp int64) (key string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(token))
	md5Ctx.Write([]byte(strconv.FormatInt(timestamp, 10)))
	data := md5Ctx.Sum(nil)
	after := hex.EncodeToString(data)
	fmt.Println(fmt.Sprintf("token=%s, timestamp=%d,data=%s", token, timestamp, after))
	return after
}

func CanonicalAddr(host string, port int) (addr string) {
	if port == 80 || port == 443 {
		addr = host
	} else {
		addr = net.JoinHostPort(host, strconv.Itoa(port))
	}
	return
}

func ParseRangeNumbers(rangeStr string) (numbers []int64, err error) {
	rangeStr = strings.TrimSpace(rangeStr)
	numbers = make([]int64, 0)
	// e.g. 1000-2000,2001,2002,3000-4000
	numRanges := strings.Split(rangeStr, ",")
	for _, numRangeStr := range numRanges {
		// 1000-2000 or 2001
		numArray := strings.Split(numRangeStr, "-")
		// length: only 1 or 2 is correct
		rangeType := len(numArray)
		switch rangeType {
		case 1:
			// single number
			singleNum, errRet := strconv.ParseInt(strings.TrimSpace(numArray[0]), 10, 64)
			if errRet != nil {
				err = fmt.Errorf("range number is invalid, %v", errRet)
				return
			}
			numbers = append(numbers, singleNum)
		case 2:
			// range numbers
			min, errRet := strconv.ParseInt(strings.TrimSpace(numArray[0]), 10, 64)
			if errRet != nil {
				err = fmt.Errorf("range number is invalid, %v", errRet)
				return
			}
			max, errRet := strconv.ParseInt(strings.TrimSpace(numArray[1]), 10, 64)
			if errRet != nil {
				err = fmt.Errorf("range number is invalid, %v", errRet)
				return
			}
			if max < min {
				err = fmt.Errorf("range number is invalid")
				return
			}
			for i := min; i <= max; i++ {
				numbers = append(numbers, i)
			}
		default:
			err = fmt.Errorf("range number is invalid")
			return
		}
	}
	return
}

func GenerateResponseErrorString(summary string, err error, detailed bool) string {
	if detailed {
		return err.Error()
	}
	return summary
}

func RandomSleep(duration time.Duration, minRatio, maxRatio float64) time.Duration {
	min := int64(minRatio * 1000.0)
	max := int64(maxRatio * 1000.0)
	var n int64
	if max <= min {
		n = min
	} else {
		n = mathrand.Int64N(max-min) + min
	}
	d := duration * time.Duration(n) / time.Duration(1000)
	time.Sleep(d)
	return d
}

func ConstantTimeEqString(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}

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

package controllers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// เข้ารหัส (ของใหม่ การ encrypt จะไม่ถูกเปลี่ยน)
func EncryptGenV2(stringToEncrypt string) (encryptedString string) {
	bKey := []byte(os.Getenv("AES_256_KEY"))
	bIV := []byte(os.Getenv("AES_256_IV"))
	bPlaintext := PKCS5PaddingV2([]byte(stringToEncrypt), aes.BlockSize, len(stringToEncrypt))
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

// เข้ารหัส สำหรับ V2
func PKCS5PaddingV2(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// ถอดรหัส (ของใหม่ การ Decrypt)
func DecryptGenV2(encryptedString string) (decryptedString string) {
	bKey := []byte(os.Getenv("AES_256_KEY"))
	bIV := []byte(os.Getenv("AES_256_IV"))
	cipherTextDecoded, err := hex.DecodeString(encryptedString)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

// แปลง Json เป็น Map
func JsonToMap(b []byte) (m map[string]interface{}) {
	json.Unmarshal(b, &m)
	return
}

func JsonToArray(b []byte) (i []interface{}) {
	json.Unmarshal(b, &i)
	return
}

// แปลง String เป็น Int
func StrToInt(str string) (i int) {
	i, _ = strconv.Atoi(str)
	return
}

// แปลง Int เป็น String
func IntToStr(i int) (str string) {
	str = strconv.Itoa(i)
	return
}

// แปลง Map เป็น  Json
func MapToJson(m interface{}) (b []byte) {
	b, _ = json.Marshal(m)
	return
}

func InterfaceToMap(b interface{}) (m map[string]interface{}) {
	bytedata, _ := json.Marshal(b)
	json.Unmarshal(bytedata, &m)

	return
}

func CheckArray(a []int, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

func CheckArrayInt64(a []int64, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

func CheckArrayString(a []string, i interface{}) (found bool) {
	for index := range a {
		if a[index] == i {
			return true
		}
	}
	return
}

// ปัดทศนิยม
func FloatPrecision(num float64, precision int) float64 {
	p := math.Pow10(precision)
	value := float64(int(num*p)) / p
	return value
}

// รับค่าได้เฉพาะอักษร
func Is_alphanum(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9ก-๛]*$`).MatchString(word)
}

// รับค่าได้ตัวเลขที่เป็น str
func Is_StrNumber(word string) bool {
	return regexp.MustCompile(`^[0-9]*$`).MatchString(word)
}

// ตัดคำ
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

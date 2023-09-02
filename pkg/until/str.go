package until

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
	"unicode"
	"unsafe"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandomString 生成随机字符串。参考：https://colobu.com/2018/09/02/generate-random-string-in-Go/
func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func IncludeLetter(str string) bool {
	runes := []rune(str)
	for _, r := range runes {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Println("r", r)
			return true
		}
	}
	return false
}

func IsDigit(str string) bool {
	for _, x := range []rune(str) {
		if !unicode.IsDigit(x) {
			return false
		}
	}
	return true
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func BytesToString(data *[]byte) string {
	// 使用地址方式转换，避免 string 在内存中重新拷贝一份
	return *(*string)(unsafe.Pointer(data))
}

func StringToBytes(data string) (b []byte) {
	// 使用地址方式转换，避免 string 在内存中重新拷贝一份
	*(*string)(unsafe.Pointer(&b)) = data
	(*reflect.SliceHeader)(unsafe.Pointer(&b)).Cap = len(data)
	return
}

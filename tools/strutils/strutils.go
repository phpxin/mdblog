package strutils

import (
	"crypto/md5"
	"encoding/hex"
)

func UcFirst(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] += - 32
	}
	return string(strArry)
}

func Md5(str string) string {
	result := md5.Sum([]byte(str))
	return hex.EncodeToString(result[:])
}
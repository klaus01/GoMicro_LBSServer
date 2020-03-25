package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// Md5 字符串 md5
func Md5(str string) string {
	if len(str) > 0 {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(str))
		cipherStr := md5Ctx.Sum(nil)
		return hex.EncodeToString(cipherStr)
	}
	return ""
}

// RandomInt 返回随机整数 [0, max]
func RandomInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(max)
}

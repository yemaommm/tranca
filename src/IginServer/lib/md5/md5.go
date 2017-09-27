package md5

import (
	"crypto/md5"
	// "fmt"
	"encoding/hex"
)

func Md5(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

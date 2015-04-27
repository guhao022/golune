package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5WithSalt(str, salt string) (enStr string) {
	h := md5.New()
	h.Write([]byte(str + salt)) // 需要加密的字符串
	enStr = hex.EncodeToString(h.Sum(nil))
	return
}

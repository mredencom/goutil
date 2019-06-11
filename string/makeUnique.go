package string

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"strconv"
)

// 制作特征值方法一
func MakeUnique(obj interface{}) string {
	b, _ := json.Marshal(obj)
	hash := fnv.New64()
	_, _ = hash.Write(b)
	return strconv.FormatUint(hash.Sum64(), 10)
}


// 制作特征值方法二
func MakeMd5(obj interface{}, length int) string {
	if length > 32 {
		length = 32
	}
	h := md5.New()
	baseString, _ := json.Marshal(obj)
	h.Write([]byte(baseString))
	s := hex.EncodeToString(h.Sum(nil))
	return s[:length]
}
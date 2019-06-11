package string

import "hash/fnv"

//hash 转字符串 和stringToHash.go 配合使用
func HashToString(encode string) uint64 {
	hash := fnv.New64()
	_, _ = hash.Write([]byte(encode))
	return hash.Sum64()
}
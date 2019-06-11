package string

import (
	"fmt"
	"hash/crc32"
)

//string to hash
//string 转 hash
func StringToHash(s string) string {
	const IEEE = 0xedb88320
	var IEEETable = crc32.MakeTable(IEEE)
	hash := fmt.Sprintf("%x", crc32.Checksum([]byte(s), IEEETable))
	return hash
}


package string

import "regexp"

// The IsNum judges string is number or not.
//检查一个字符串是否是数字
func IsNum(a string) bool {
	reg, _ := regexp.Compile("^\\d+$")
	return reg.MatchString(a)
}

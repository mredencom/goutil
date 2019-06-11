package file

import "strings"

// 将文件名非法字符替换为相似字符
func FileNameReplace(fileName string) string {
	var q = 1
	r := []rune(fileName)
	size := len(r)
	for i := 0; i < size; i++ {
		switch r[i] {
		case '"':
			if q%2 == 1 {
				r[i] = '“'
			} else {
				r[i] = '”'
			}
			q++
		case ':':
			r[i] = '：'
		case '*':
			r[i] = '×'
		case '<':
			r[i] = '＜'
		case '>':
			r[i] = '＞'
		case '?':
			r[i] = '？'
		case '/':
			r[i] = '／'
		case '|':
			r[i] = '∣'
		case '\\':
			r[i] = '╲'
		}
	}
	return strings.Replace(string(r), USE_KEYIN, ``, -1)
}
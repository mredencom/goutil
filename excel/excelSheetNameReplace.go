package excel

import "strings"

// 将Excel工作表名中非法字符替换为下划线
func ExcelSheetNameReplace(fileName string) string {
	r := []rune(fileName)
	size := len(r)
	for i := 0; i < size; i++ {
		switch r[i] {
		case ':', '：', '*', '?', '？', '/', '／', '\\', '╲', ']', '[':
			r[i] = '_'
		}
	}
	return strings.Replace(string(r), USE_KEYIN, ``, -1)
}
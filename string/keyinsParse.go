package string

import "strings"

// 切分用户输入的自定义信息
func KeyinsParse(keyins string) []string {
	keyins = strings.TrimSpace(keyins)
	if keyins == "" {
		return []string{}
	}
	for _, v := range re.FindAllString(keyins, -1) {
		keyins = strings.Replace(keyins, v, "><", -1)
	}
	m := map[string]bool{}
	for _, v := range strings.Split(keyins, "><") {
		v = strings.TrimPrefix(v, "<")
		v = strings.TrimSuffix(v, ">")
		if v == "" {
			continue
		}
		m[v] = true
	}
	s := make([]string, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	return s
}
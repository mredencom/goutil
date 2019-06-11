package json

import "encoding/json"

//json转换成map
func JsonToMap(jsonStr string) map[string]string {
	cookieMap := map[string]string{}
	_ = json.Unmarshal([]byte(jsonStr), &cookieMap)
	return cookieMap
}

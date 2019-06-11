package xml

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"
	"strings"
)

// simple xml to string  support utf8
//把xml文件转换为map 只是简单的xml
func XML2mapstr(xmldoc string) map[string]string {
	var t xml.Token
	var err error
	inputReader := strings.NewReader(xmldoc)
	decoder := xml.NewDecoder(inputReader)
	decoder.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	m := make(map[string]string, 32)
	key := ""
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			key = token.Name.Local
		case xml.CharData:
			content := Bytes2String([]byte(token))
			m[key] = content
		default:
			// ...
		}
	}
	return m
}
package http

import "net/url"

//拼接url参数
func UrlQuery(finalUrl *url.URL, querys map[string]string) {
	params := url.Values{}
	for k, v := range querys {
		params.Set(k, v)
	}
	finalUrl.RawQuery = params.Encode()
}

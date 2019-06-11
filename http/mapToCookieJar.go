package http

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)
//把map中的值设置到cookie中去
func MapToCookieJar(cookieMap map[string]string, domain string, finalUrl *url.URL) http.CookieJar {
	var cookies []*http.Cookie
	for cookieName, cookieValue := range cookieMap {
		cookie := &http.Cookie{
			Name:   cookieName,
			Value:  cookieValue,
			Path:   "/",
			Domain: domain,
		}
		cookies = append(cookies, cookie)
	}
	cookieJar, _ := cookiejar.New(nil)
	cookieJar.SetCookies(finalUrl, cookies)
	return cookieJar
}
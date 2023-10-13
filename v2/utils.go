package v2

import (
	"net/http"
	"net/url"

	"golang.org/x/net/http/httpproxy"
)

func proxyFromUrl(proxyUrl string) func(*http.Request) (*url.URL, error) {
	return func(req *http.Request) (*url.URL, error) {
		cnf := &httpproxy.Config{
			HTTPProxy:  proxyUrl,
			HTTPSProxy: proxyUrl,
			NoProxy:    "",
			CGI:        false,
		}
		return cnf.ProxyFunc()(req.URL)
	}
}

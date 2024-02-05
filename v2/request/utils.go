package request

import (
	"bytes"
	"encoding/json"
	"golang.org/x/net/http/httpproxy"
	"net/http"
	"net/url"

	"github.com/YuzuWiki/Pixivlee/v2"
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

func encodeURL(u string, data v2.Query) (string, error) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if data != nil {
		URL.RawQuery = data.Encode()
	}

	return URL.String(), nil
}

func encodeBody(params v2.Params) *bytes.Buffer {
	if body, err := json.Marshal(params); err == nil {
		return bytes.NewBuffer(body)
	}
	return nil
}

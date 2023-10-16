package v2

import (
	"bytes"
	"encoding/json"
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

func encodeURL(u string, data Query) (string, error) {
	URL, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	if data != nil {
		URL.RawQuery = data.Encode()
	}

	return URL.String(), nil
}

func encodeBody(params Params) *bytes.Buffer {
	if body, err := json.Marshal(params); err == nil {
		return bytes.NewBuffer(body)
	}
	return nil
}

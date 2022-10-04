package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/http/httpproxy"

	"github.com/YuzuWiki/Pixivlee/common"
)

type transport struct {
	transport http.Transport

	// mutex
	mu sync.Mutex
}

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

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return t.transport.RoundTrip(req)
}

func (t *transport) SetProxy(proxyUrl string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(proxyUrl) > 0 && t.transport.Proxy == nil {
		t.transport.Proxy = proxyFromUrl(proxyUrl)
		return nil
	}
	return fmt.Errorf("SetProxy Error")
}

func (t *transport) UnSetProxy() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.transport.Proxy = nil
	return nil
}

func NewTransport() common.ITransport {
	return &transport{
		mu: sync.Mutex{},

		transport: http.Transport{
			DisableKeepAlives: true,
			Proxy:             nil,
		},
	}
}

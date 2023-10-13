package v2

import (
	"fmt"
	"net/http"
	"sync"
)

type transport struct {
	t http.Transport

	mu sync.Mutex
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	return t.t.RoundTrip(req)
}

func (t *transport) SetProxy(proxyUrl string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(proxyUrl) > 0 && t.t.Proxy == nil {
		t.t.Proxy = proxyFromUrl(proxyUrl)
		return nil
	}
	return fmt.Errorf("proxy setup error")
}

func (t *transport) UnSetProxy() error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.t.Proxy = nil
	return nil
}

func NewTransport() ITransport {
	return &transport{
		mu: sync.Mutex{},

		t: http.Transport{
			DisableKeepAlives: true,
			Proxy:             nil,
		},
	}
}

type request struct {
	http.Client

	t transport

	beforeHooks []BeforeHook
	afterHooks  []AfterHook
}

func (r *request) do() {
}

func (r *request) Post() {
}

func (r *request) Get() {

}

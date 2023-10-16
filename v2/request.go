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

func newRequest(method, u string, query Query, params Params) (*http.Request, error) {
	u, err := encodeURL(u, query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u, encodeBody(params))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func doHooks[T *http.Request | *http.Response](hooks *[]func(T) error, body T) (err error) {
	funcs := *hooks
	for idx := range funcs {
		if err = funcs[idx](body); err != nil {
			return
		}
	}
	return
}

func (r *request) do(method, u string, query Query, params Params) (resp *http.Response, err error) {
	req, err := newRequest(method, u, query, params)
	if err != nil {
		return nil, err
	}

	if err = doHooks[*http.Request](&r.beforeHooks, req); err != nil {
		return nil, err
	}

	switch req.URL.Hostname() {
	case "www.fanbox.cc":
		resp, err = r.Transport.RoundTrip(req)
	default:
		resp, err = r.Client.Do(req)
	}

	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
	} else {
		if err = doHooks[*http.Response](&r.afterHooks, resp); err != nil {
			return nil, err
		}
	}

	return resp, err
}

func (r *request) Post() {
}

func (r *request) Get() {

}

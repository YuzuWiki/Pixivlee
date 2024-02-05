package request

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee/v2"
)

type request struct {
	http.Client

	t v2.ITransport

	beforeHooks []func(req *http.Request) error
	afterHooks  []func(resp *http.Response) error
}

func newRequest(method, u string, query v2.Query, params v2.Params) (*http.Request, error) {
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
			return err
		}
	}
	return
}

func (r *request) do(method, u string, query v2.Query, params v2.Params) (resp *http.Response, err error) {
	req, err := newRequest(method, u, query, params)
	if err != nil {
		return nil, err
	}

	if err = doHooks[*http.Request](&r.beforeHooks, req); err != nil {
		return nil, err
	}

	switch req.URL.Hostname() {
	case v2.HostNameFanbox:
		resp, err = r.Client.Do(req)

	case v2.HostNamePixiv:
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

func (r *request) Head(u string, query v2.Query, params v2.Params) (*http.Response, error) {
	return r.do(http.MethodHead, u, query, params)
}

func (r *request) Get(u string, query v2.Query, params v2.Params) (*http.Response, error) {
	return r.do(http.MethodGet, u, query, params)
}

func (r *request) Post(u string, query v2.Query, params v2.Params) (*http.Response, error) {
	return r.do(http.MethodPost, u, query, params)
}

func (r *request) Put(u string, query v2.Query, params v2.Params) (*http.Response, error) {
	return r.do(http.MethodPut, u, query, params)
}

func (r *request) Delete(u string, query v2.Query, params v2.Params) (*http.Response, error) {
	return r.do(http.MethodDelete, u, query, params)
}

func (r *request) SetProxy(proxyUrl string) error {
	return r.t.SetProxy(proxyUrl)
}

func (r *request) UnSetProxy() error {
	return r.t.UnSetProxy()

}

func NewRequest() v2.IRequest {
	t := NewTransport()
	return &request{
		Client: http.Client{Transport: t},
		t:      t,
	}
}

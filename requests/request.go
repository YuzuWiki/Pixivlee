package requests

import (
	"net/http"
	"strings"

	"github.com/YuzuWiki/Pixivlee/common"
)

type requests struct {
	http.Client

	// Transport
	Transport common.ITransport

	// 钩子函数..
	beforeHooks []common.BeforeHook
	afterHooks  []common.AfterHook
}

func doHooks[T *http.Request | *http.Response](hooks []func(T) error, body T) (err error) {
	for idx := range hooks {
		if err = hooks[idx](body); err != nil {
			return
		}
	}
	return
}

func newRequest(u, method string, query *common.Query, params *common.Params) (req *http.Request, err error) {
	if u, err = common.EncodeURL(u, query); err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(method, u, common.EncodeBody(params)); err != nil {
		return nil, err
	}
	return req, nil
}

func (r *requests) do(method, u string, query *common.Query, params *common.Params) (resp *http.Response, err error) {
	var req *http.Request
	if req, err = newRequest(method, u, query, params); err != nil {
		return
	}

	if err = doHooks(r.beforeHooks, req); err != nil {
		return
	}

	// https://www.fanbox.cc/
	switch strings.SplitN(strings.Replace(req.URL.Path, "/", "", 1), "/", 3)[0] {
	case "fanbox":
		resp, err = r.Transport.RoundTrip(req)
	default:
		resp, err = r.Client.Do(req)
	}

	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return nil, err
	}

	if err = doHooks(r.afterHooks, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *requests) Head(u string, query *common.Query, params *common.Params) (*http.Response, error) {
	return r.do(u, http.MethodHead, query, params)
}

func (r *requests) Get(u string, query *common.Query, params *common.Params) (*http.Response, error) {
	return r.do(u, http.MethodGet, query, params)
}

func (r *requests) Post(u string, query *common.Query, params *common.Params) (*http.Response, error) {
	return r.do(u, http.MethodPost, query, params)
}

func (r *requests) Put(u string, query *common.Query, params *common.Params) (*http.Response, error) {
	return r.do(u, http.MethodPut, query, params)
}

func (r *requests) Delete(u string, query *common.Query, params *common.Params) (*http.Response, error) {
	return r.do(u, http.MethodDelete, query, params)
}

func (r *requests) SetProxy(proxyUrl string) error {
	return r.Transport.SetProxy(proxyUrl)
}

func (r *requests) UnSetProxy() error {
	return r.Transport.UnSetProxy()
}

func (r *requests) BeforeHooks(fns ...common.BeforeHook) {
	for idx := range fns {
		r.beforeHooks = append(r.beforeHooks, fns[idx])
	}
}

func (r *requests) AfterHooks(fns ...common.AfterHook) {
	for idx := range fns {
		r.afterHooks = append(r.afterHooks, fns[idx])
	}
}

func NewRequest() common.IRequest {
	t := NewTransport()
	return &requests{
		Client:    http.Client{Transport: t},
		Transport: t,
	}
}

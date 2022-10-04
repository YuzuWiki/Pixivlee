package common

import (
	"net/http"
	"net/url"
)

type (
	Query        = url.Values          // requests data
	Params       = map[string]struct{} // requests body
	BeforeHook   = func(req *http.Request) error
	AfterHook    = func(resp *http.Response) error
	HeaderOption struct {
		Key   string
		Value string
	}
)

type ITransport interface {
	http.RoundTripper

	SetProxy(string) error
	UnSetProxy() error
}

type IClient interface {
	Head(string, *Query, *Params) (*http.Response, error)
	Get(string, *Query, *Params) (*http.Response, error)
	Post(string, *Query, *Params) (*http.Response, error)
	Put(string, *Query, *Params) (*http.Response, error)
	Delete(string, *Query, *Params) (*http.Response, error)

	SetProxy(string) error
	UnSetProxy() error
}

type IRequest interface {
	IClient

	SetCookies(string, ...*http.Cookie) error
	SetHeader(...HeaderOption)

	BeforeHooks(...BeforeHook)
	AfterHooks(...AfterHook)
}

type IContext interface {
	PhpSessID() string
}

type IPool interface {
	Get(string) (IClient, error)
	Del(string)
}

package v2

import (
	"net/http"
	"net/url"
)

type (
	TPid        uint64 // pid type(alise)
	TArtId      uint64
	TIllustType uint8

	Query        = url.Values
	Params       = map[string]struct{}
	HeaderOption struct {
		Key   string
		Value string
	}
)

// IPixiver pixiver
type IPixiver interface {
	Pid() TPid
	SessID() string
	IsEnable() bool
}

// IPool pixiver pool
type IPool interface {
	Push(pixiver IPixiver) error
	Pop() IPixiver
	List() []IPixiver
}

type ITransport interface {
	http.RoundTripper

	SetProxy(string) error
	UnSetProxy() error
}

type IRequest interface {
	Head(u string, query Query, params Params) (*http.Response, error)
	Get(u string, query Query, params Params) (*http.Response, error)
	Post(u string, query Query, params Params) (*http.Response, error)
	Put(u string, query Query, params Params) (*http.Response, error)
	Delete(u string, query Query, params Params) (*http.Response, error)
}

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
	BeforeHook   = func(req *http.Request) error
	AfterHook    = func(resp *http.Response) error
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

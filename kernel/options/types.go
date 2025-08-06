package options

import (
	"github.com/YuzuWiki/Pixivlee/kernel/httpx"
)

type Middleware = func() error

type HttpOption struct {
	IsDebug  bool
	Timeout  int
	ProxyUrl string

	BaseURl   string
	Host      string
	UserAgent string

	OnBeforeRequest []httpx.RequestMiddleware
	OnAfterResponse []httpx.ResponseMiddleware
}

package kernel

import (
	resty "github.com/go-resty/resty/v2"
)

type HttpOption struct {
	IsDebug   bool
	Timeout   int
	ProxyUrl  string
	UserAgent string

	OnBeforeRequest []resty.RequestMiddleware
	OnAfterResponse []resty.ResponseMiddleware
}

type PixivOption struct {
	PhpSessId string
}

type Options struct {
	Http  HttpOption
	Pixiv PixivOption
}

package kernel

import "github.com/imroc/req/v3"

type HttpOption struct {
	IsDebug   bool
	Timeout   int
	ProxyUrl  string
	UserAgent string

	OnBeforeRequest []req.RequestMiddleware
	OnAfterResponse []req.ResponseMiddleware
}

type PixivOption struct {
	PhpSessId string
}

type Options struct {
	Http  HttpOption
	Pixiv PixivOption
}

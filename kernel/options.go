package kernel

import (
	"github.com/YuzuWiki/Pixivlee/kernel/request"
)

type HttpOption struct {
	IsDebug   bool
	Timeout   int
	ProxyUrl  string
	UserAgent string

	OnBeforeRequest []request.RequestMiddleware
	OnAfterResponse []request.ResponseMiddleware
}

type PixivOption struct {
	PhpSessId string
}

type Options struct {
	Http  HttpOption
	Pixiv PixivOption
}

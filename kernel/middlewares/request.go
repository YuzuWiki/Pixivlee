package middlewares

import (
	"strings"

	"github.com/YuzuWiki/Pixivlee/kernel/httpx"
)

const (
	DEFAULT_USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0"
)

func DefaultRequest(ua, referer, host string) func(iClient httpx.IClient, iRequest httpx.IRequest) error {
	ua = strings.TrimSpace(ua)
	if ua == "" {
		ua = DEFAULT_USER_AGENT
	}

	return func(iClient httpx.IClient, iRequest httpx.IRequest) error {
		iRequest.SetHeader("User-Agent", ua).
			SetHeader("referer", strings.TrimSpace(referer)).
			SetHeader("origin", strings.TrimSpace(host))
		return nil
	}
}

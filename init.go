package Pixivlee

import (
	"github.com/imroc/req/v3"
	"time"
)

var (
	requestClient *req.Client = nil
)

func defaultHeader(c *req.Client, host string) *req.Client {
	c.SetBaseURL("https://" + host)

	c.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		req.SetHeader("User-Agent", UserAgent).
			SetHeader("referer", "https://"+host)
		return nil
	})

	return c
}

func SetProxy(proxyUrl string) {
	requestClient.SetProxyURL(proxyUrl)
}

func SetTimeOut(second time.Duration) {
	requestClient.SetTimeout(second * time.Second)
}

func init() {
	requestClient = defaultHeader(req.C(), PixivHost).SetTimeout(15 * time.Second).EnableDebugLog()
}

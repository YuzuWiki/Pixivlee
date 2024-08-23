package Pixivlee

import (
	"time"

	"github.com/imroc/req/v3"
)

var (
	requests *req.Client = nil
)

func SetProxy(proxyUrl string) *req.Client {
	return requests.SetProxyURL(proxyUrl)
}

func UnSetProxy() *req.Client {
	requests.SetProxy(nil)

	// close idle connect
	//requests.CloseIdleConnections()
	return requests
}

func SetTimeOut(second int) *req.Client {
	return requests.SetTimeout(time.Duration(second) * time.Second)
}

func OnBeforeRequest(fn req.RequestMiddleware) *req.Client {
	return requests.OnBeforeRequest(fn)
}

func OnAfterResponse(fn req.ResponseMiddleware) *req.Client {
	return requests.OnAfterResponse(fn)
}

func EnableDebug() *req.Client {
	return requests.EnableDebugLog()
}

func init() {
	requests = req.C()

	// default base url
	requests.SetBaseURL("https://" + PIXIV_HOST)

	// default header
	requests.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
		req.SetHeader("User-Agent", USER_AGENT).
			SetHeader("referer", "https://"+PIXIV_HOST)
		return nil
	})

	// default timeout
	requests.SetTimeout(15 * time.Second)
}

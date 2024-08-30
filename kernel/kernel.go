package kernel

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Kernel struct {
	pixiver types.IPixiver

	requests *req.Client
}

func (k *Kernel) SetPixiver(pixiver types.IPixiver) {
	k.pixiver = pixiver
}

func (k *Kernel) EnableDebug() *Kernel {
	k.requests.EnableDebugLog()
	return k
}

func (k *Kernel) SetBaseURL(baseUrl string) *Kernel {
	k.requests.SetBaseURL(baseUrl)
	return k
}

func (k *Kernel) SetProxy(proxyUrl string) *Kernel {
	k.requests.SetProxyURL(proxyUrl)
	return k
}

func (k *Kernel) UnSetProxy() *Kernel {
	k.requests.SetProxy(nil)
	return k
}

func (k *Kernel) SetTimeOut(second int) *Kernel {
	k.requests.SetTimeout(time.Duration(second) * time.Second)
	return k
}

func (k *Kernel) OnBeforeRequest(fn req.RequestMiddleware) *Kernel {
	k.requests.OnBeforeRequest(fn)
	return k
}

func (k *Kernel) OnAfterResponse(fn req.ResponseMiddleware) *Kernel {
	k.requests.OnAfterResponse(fn)
	return k
}

func (k *Kernel) NewRequests() *req.Request {
	r := k.requests.NewRequest()

	// set cookie
	r.SetCookies(
		&http.Cookie{
			Name:   "PHPSESSID",
			Value:  k.pixiver.SessionID(),
			Path:   "/",
			Domain: ".pixiv.net",
		})

	// default params
	r.AddQueryParam("lang", "jp")
	return r
}

func NewKernel() *Kernel {
	return &Kernel{
		pixiver:  nil,
		requests: req.NewClient(),
	}
}

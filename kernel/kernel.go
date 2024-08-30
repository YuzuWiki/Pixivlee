package kernel

import (
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Kernel struct {
	pixiver types.IPixiver

	session *resty.Client
}

func (k *Kernel) SetPixiver(pixiver types.IPixiver) {
	k.pixiver = pixiver
}

func (k *Kernel) EnableDebug() *Kernel {
	k.session.Debug = true
	return k
}

func (k *Kernel) SetBaseURL(baseUrl string) *Kernel {
	k.session.SetBaseURL(baseUrl)
	return k
}

func (k *Kernel) SetProxy(proxyUrl string) *Kernel {
	k.session.SetProxy(proxyUrl)
	return k
}

func (k *Kernel) UnSetProxy() *Kernel {
	k.session.RemoveProxy()
	return k
}

func (k *Kernel) SetTimeOut(second int) *Kernel {
	k.session.SetTimeout(time.Duration(second) * time.Second)
	return k
}

func (k *Kernel) OnBeforeRequest(fn resty.RequestMiddleware) *Kernel {
	k.session.OnBeforeRequest(fn)
	return k
}

func (k *Kernel) OnAfterResponse(fn resty.ResponseMiddleware) *Kernel {
	k.session.OnAfterResponse(fn)
	return k
}

func (k *Kernel) NewRequests() *resty.Request {
	r := k.session.NewRequest()

	// set cookie
	r.SetCookies([]*http.Cookie{
		{
			Name:   "PHPSESSID",
			Value:  k.pixiver.SessionID(),
			Path:   "/",
			Domain: ".pixiv.net",
		},
	})

	// default params
	r.SetQueryParam("lang", "jp")
	return r
}

func NewKernel() *Kernel {
	return &Kernel{
		pixiver: nil,
		session: resty.New(),
	}
}

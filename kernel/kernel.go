package kernel

import (
	"net/http"
	"time"

	"github.com/YuzuWiki/Pixivlee/kernel/request"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Kernel struct {
	pixiver types.IPixiver

	session request.ISession
}

func (k *Kernel) SetPixiver(pixiver types.IPixiver) {
	k.pixiver = pixiver
}

func (k *Kernel) EnableDebug() *Kernel {
	k.session.SetDebug(true)
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

func (k *Kernel) OnBeforeRequest(fn request.RequestMiddleware) *Kernel {
	k.session.OnBeforeRequest(fn)
	return k
}

func (k *Kernel) OnAfterResponse(fn request.ResponseMiddleware) *Kernel {
	k.session.OnAfterResponse(fn)
	return k
}

func (k *Kernel) NewRequests() request.IRequest {
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
		session: request.New(),
	}
}

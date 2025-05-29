package kernel

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee/kernel/httpx"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Kernel struct {
	account types.IAccount

	session httpx.ISession
}

func (k *Kernel) Session() httpx.ISession {
	return k.session
}

func (k *Kernel) Account() types.IAccount {
	return k.account
}

func (k *Kernel) NewRequests() httpx.IRequest {
	r := k.session.NewRequest()
	// fixme: 职责划分错误
	// set cookie
	r.SetCookies([]*http.Cookie{
		{

			Name:   "PHPSESSID",
			Value:  k.Account().SessionID(),
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
		account: nil,
		session: httpx.New(),
	}
}

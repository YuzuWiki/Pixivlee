package types

import (
	http2 "github.com/YuzuWiki/Pixivlee/kernel/httpx"
	"net/http"
)

type IAccount interface {
	IPixiver

	Cookies() []*http.Cookie
}

type IKernel interface {
	Account() IAccount
	Session() http2.ISession

	NewRequests() http2.IRequest
}

type IPixiver interface {
	Pid() TPid
	SessionID() string
}

package types

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee/kernel/httpx"
)

type IAccount interface {
	IPixiver

	Cookies() []*http.Cookie
}

type IKernel interface {
	Account() IAccount
	Session() httpx.IClient

	NewRequests() httpx.IRequest
}

type IPixiver interface {
	Pid() TPid
	SessionID() string
}

package kernel

import (
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/YuzuWiki/Pixivlee/kernel/internal/token"
)

type TRequest = *resty.Client

const (
	DEFAULT_USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0"
	DEFAULT_LANG       = "jp"
	DOMAIN_PIXIV       = ".pixiv.net"
	HOST_PIXIV         = "www.pixiv.net"
	SESSION_ID_PIXIV   = token.PIXIV_SESSION_ID
	DOMAIN_FANBOX      = ".fanbox.cc"
	HOST_FANBOX        = "www.fanbox.cc"
	SESSION_ID_FANBOX  = token.FANBOX_SESSION_ID
)

type Kernel struct {
	c       *resty.Client  // http client
	token   *token.Token   // pixiver token
	cookies []*http.Cookie // session cookie
}

func (k *Kernel) C() *resty.Client {
	return k.c
}

// R new request
func (k *Kernel) R() *resty.Request {
	r := k.c.R()

	// set pixiv
	if len(k.token.SidPixiv) > 0 {
		r.SetCookie(&http.Cookie{
			Name:   SESSION_ID_PIXIV,
			Value:  SESSION_ID_PIXIV + "=" + k.token.SidPixiv,
			Path:   "/",
			Domain: DOMAIN_PIXIV,
		})

		//r.SetHeader("origin", "https://www.pixiv.net")

		//r.SetHeader("referer", "https://www.pixiv.net/")
	}

	// set fanbox
	if len(k.token.SidFanbox) > 0 {
		r.SetCookie(&http.Cookie{
			Name:   SESSION_ID_FANBOX,
			Value:  SESSION_ID_FANBOX + "=" + k.token.SidFanbox,
			Path:   "/",
			Domain: DOMAIN_FANBOX,
		})

		r.SetHeader("origin", "https://www.fanbox.cc")

		r.SetHeader("referer", "https://www.fanbox.cc/")
	}

	// default User-Agent
	r.SetHeader("User-Agent", DEFAULT_USER_AGENT)

	// default params
	r.SetQueryParam("lang", DEFAULT_LANG)
	return r
}

func NewKernel(cookies string) (*Kernel, error) {
	t, err := token.New(cookies)
	if err != nil {
		return nil, err
	}

	return &Kernel{
		token: t,
		c:     resty.New(),
	}, nil
}

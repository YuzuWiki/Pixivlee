package request

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type TSession struct {
	r *resty.Client
}

func (c *TSession) SetDebug(d bool) ISession {
	c.r.SetDebug(d)
	return c
}

func (c *TSession) SetLogger(logger ILogger) ISession {
	c.r.SetLogger(logger)
	return c
}

func (c *TSession) SetTimeout(timeout time.Duration) ISession {
	c.r.SetTimeout(timeout)
	return c
}

func (c *TSession) SetProxy(url string) ISession {
	c.r.SetProxy(url)
	return c
}

func (c *TSession) RemoveProxy() ISession {
	c.r.RemoveProxy()
	return c
}

func (c *TSession) SetBaseURL(url string) ISession {
	c.r.SetBaseURL(url)
	return c
}

func (c *TSession) SetHeader(header, value string) ISession {
	c.r.SetHeader(header, value)
	return c
}

func (c *TSession) SetHeaders(headers map[string]string) ISession {
	c.r.SetHeaders(headers)
	return c
}

func (c *TSession) SetCookie(r *http.Cookie) ISession {
	c.r.SetCookie(r)
	return c
}
func (c *TSession) SetCookies(rs []*http.Cookie) ISession {
	c.r.SetCookies(rs)
	return c
}

func (c *TSession) SetQueryParam(param, value string) ISession {
	c.r.SetQueryParam(param, value)
	return c
}

func (c *TSession) SetQueryParams(params map[string]string) ISession {
	c.r.SetPathParams(params)
	return c
}

func (c *TSession) NewRequest() IRequest {
	return &TRequest{req: c.r.NewRequest()}
}

func (c *TSession) OnBeforeRequest(fn RequestMiddleware) ISession {
	c.r.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		return fn(&TSession{r: client}, &TRequest{req: request})
	})
	return c
}

func (c *TSession) OnAfterResponse(fn ResponseMiddleware) ISession {
	c.r.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		return fn(&TSession{r: client}, &TResponse{response})
	})
	return c
}

func (c *TSession) OnError(fn ErrorHook) ISession {
	c.r.OnError(func(request *resty.Request, err error) {
		fn(&TRequest{req: request}, err)
	})
	return c
}

func (c *TSession) OnSuccess(fn SuccessHook) ISession {
	c.r.OnSuccess(func(client *resty.Client, response *resty.Response) {
		fn(&TSession{r: client}, &TResponse{response})
	})
	return c
}

func New() ISession {
	return &TSession{r: resty.New()}
}

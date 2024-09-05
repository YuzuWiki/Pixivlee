package request

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type TClient struct {
	r *resty.Client
}

func (c *TClient) SetDebug(d bool) ISession {
	c.r.SetDebug(d)
	return c
}

func (c *TClient) SetLogger(logger ILogger) ISession {
	c.r.SetLogger(logger)
	return c
}

func (c *TClient) SetTimeout(timeout time.Duration) ISession {
	c.r.SetTimeout(timeout)
	return c
}

func (c *TClient) SetProxy(url string) ISession {
	c.r.SetProxy(url)
	return c
}

func (c *TClient) RemoveProxy() ISession {
	c.r.RemoveProxy()
	return c
}

func (c *TClient) SetBaseURL(url string) ISession {
	c.r.SetBaseURL(url)
	return c
}

func (c *TClient) SetHeader(header, value string) ISession {
	c.r.SetHeader(header, value)
	return c
}

func (c *TClient) SetHeaders(headers map[string]string) ISession {
	c.r.SetHeaders(headers)
	return c
}

func (c *TClient) SetCookie(r *http.Cookie) ISession {
	c.r.SetCookie(r)
	return c
}
func (c *TClient) SetCookies(rs []*http.Cookie) ISession {
	c.r.SetCookies(rs)
	return c
}

func (c *TClient) SetQueryParam(param, value string) ISession {
	c.r.SetQueryParam(param, value)
	return c
}

func (c *TClient) SetQueryParams(params map[string]string) ISession {
	c.r.SetPathParams(params)
	return c
}

func (c *TClient) NewRequest() IRequest {
	return &TRequest{req: c.r.NewRequest()}
}

func (c *TClient) OnBeforeRequest(fn RequestMiddleware) ISession {
	c.r.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		return fn(&TClient{r: client}, &TRequest{req: request})
	})
	return c
}

func (c *TClient) OnAfterResponse(fn ResponseMiddleware) ISession {
	c.r.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		return fn(&TClient{r: client}, &TResponse{response})
	})
	return c
}

func (c *TClient) OnError(fn ErrorHook) ISession {
	c.r.OnError(func(request *resty.Request, err error) {
		fn(&TRequest{req: request}, err)
	})
	return c
}

func (c *TClient) OnSuccess(fn SuccessHook) ISession {
	c.r.OnSuccess(func(client *resty.Client, response *resty.Response) {
		fn(&TClient{r: client}, &TResponse{response})
	})
	return c
}

func New() ISession {
	return &TClient{r: resty.New()}
}

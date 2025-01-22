package httpx

import (
	"net/http"
	"time"
)

type ILogger interface {
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

type IResponse interface {
	IsError() bool

	StatusCode() int

	Body() []byte
	String() string

	Cookies() []*http.Cookie

	RawResponse() *http.Response
}

type IRequest interface {
	SetHeader(header, value string) IRequest
	SetHeaders(headers map[string]string) IRequest

	SetQueryParam(param, value string) IRequest
	SetQueryParams(params map[string]string) IRequest

	SetPathParam(param, value string) IRequest
	SetPathParams(params map[string]string) IRequest

	SetCookie(hc *http.Cookie) IRequest
	SetCookies(rs []*http.Cookie) IRequest

	SetBody(body any) IRequest
	SetResult(res any) IRequest

	Options(url string) (IResponse, error)
	Head(url string) (IResponse, error)
	Get(url string) (IResponse, error)
	Post(url string) (IResponse, error)
	Put(url string) (IResponse, error)
	Delete(url string) (IResponse, error)
}

type ISession interface {
	SetDebug(d bool) ISession
	SetLogger(logger ILogger) ISession
	SetTimeout(timeout time.Duration) ISession

	SetProxy(url string) ISession
	RemoveProxy() ISession

	SetBaseURL(url string) ISession

	SetHeader(header, value string) ISession
	SetHeaders(headers map[string]string) ISession

	SetCookie(r *http.Cookie) ISession
	SetCookies(rs []*http.Cookie) ISession

	SetQueryParam(param, value string) ISession
	SetQueryParams(params map[string]string) ISession

	NewRequest() IRequest

	OnBeforeRequest(RequestMiddleware) ISession
	OnAfterResponse(ResponseMiddleware) ISession

	OnError(ErrorHook) ISession
	OnSuccess(SuccessHook) ISession
}

type (
	RequestMiddleware  func(ISession, IRequest) error
	ResponseMiddleware func(ISession, IResponse) error

	ErrorHook   func(IRequest, error)
	SuccessHook func(ISession, IResponse)
)

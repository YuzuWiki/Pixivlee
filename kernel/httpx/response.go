package httpx

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

type TResponse struct {
	resp *resty.Response
}

func (r *TResponse) IsError() bool {
	return r.resp.IsError()
}

func (r *TResponse) StatusCode() int {
	return r.resp.StatusCode()
}
func (r *TResponse) Body() []byte {
	return r.resp.Body()
}
func (r *TResponse) String() string {
	return r.resp.String()
}

func (r *TResponse) Cookies() []*http.Cookie {
	return r.resp.Cookies()
}

func (r *TResponse) RawResponse() *http.Response {
	return r.resp.RawResponse
}

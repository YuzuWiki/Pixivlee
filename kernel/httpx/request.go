package httpx

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type TRequest struct {
	req *resty.Request
}

func (r *TRequest) SetHeader(header, value string) IRequest {
	r.req.SetHeader(header, value)
	return r
}

func (r *TRequest) SetHeaders(headers map[string]string) IRequest {
	r.req.SetHeaders(headers)
	return r
}

func (r *TRequest) SetQueryParam(param, value string) IRequest {
	r.req.SetQueryParam(param, value)
	return r
}

func (r *TRequest) SetQueryParams(params map[string]string) IRequest {
	r.req.SetQueryParams(params)
	return r
}

func (r *TRequest) SetPathParam(param, value string) IRequest {
	r.req.SetPathParam(param, value)
	return r
}

func (r *TRequest) SetPathParams(params map[string]string) IRequest {
	r.req.SetPathParams(params)
	return r
}

func (r *TRequest) SetCookie(hc *http.Cookie) IRequest {
	r.req.SetCookie(hc)
	return r
}

func (r *TRequest) SetCookies(rs []*http.Cookie) IRequest {
	r.req.SetCookies(rs)
	return r
}

func (r *TRequest) SetBody(body any) IRequest {
	r.req.SetBody(body)
	return r
}

func (r *TRequest) SetResult(res any) IRequest {
	r.req.SetResult(res)
	return r
}

func (r *TRequest) Options(url string) (IResponse, error) {
	resp, err := r.req.Options(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

func (r *TRequest) Head(url string) (IResponse, error) {
	resp, err := r.req.Head(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

func (r *TRequest) Get(url string) (IResponse, error) {
	resp, err := r.req.Get(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

func (r *TRequest) Post(url string) (IResponse, error) {
	resp, err := r.req.Post(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

func (r *TRequest) Put(url string) (IResponse, error) {
	resp, err := r.req.Put(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

func (r *TRequest) Delete(url string) (IResponse, error) {
	resp, err := r.req.Delete(url)
	if err != nil {
		return nil, err
	}
	return &TResponse{resp}, nil
}

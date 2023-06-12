package Pixivlee

import (
	"net/http"
	"net/url"
)

/*
  user_pool: { IPixiver{}... }
		|
		V
    context(user, proxy)
		|
		V
	request: {api, ctx, query, body}
		before_do:  user检测(IsActive, limit)
		after_do:
*/

// type alise
type (
	// TQuery request.data
	TQuery = url.Values

	// TParams request.body
	TParams = map[string]struct{}
)

// IPixiver 访问凭据
type IPixiver interface {
	Pid() string

	SessID() string
	IsActive() bool
}

// IContext 上下文信息
type IContext interface {
	// ProxyUri transport use
	ProxyUri() string
}

// IApi api实例配置配置信息
type IApi interface {
	Method() string
	Url() string
}

// IRequest request interface
type IRequest interface {
	Head(IApi, IContext, *TQuery, *TParams) (*http.Response, error)
	Get(IApi, IContext, *TQuery, *TParams) (*http.Response, error)
	Post(IApi, IContext, *TQuery, *TParams) (*http.Response, error)
	Put(IApi, IContext, *TQuery, *TParams) (*http.Response, error)
	Delete(IApi, IContext, *TQuery, *TParams) (*http.Response, error)
}

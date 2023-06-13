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
					|
					V
				client_pool: {IClient{}, ...}
			|
			V
		before_do:  user检测(IsActive, limit)
		after_do:   notify
*/

// type alise
type (
	// TQuery request.data
	TQuery = url.Values

	// TJson request.body
	TJson = map[string]struct{}

	// THeader request.header.set
	THeader = struct {
		Key   string
		Value string
	}

	// TCookie request.cookie body
	TCookie = http.Cookie
)

// IPixiver 访问凭据
type IPixiver interface {
	Pid() string
	SessID() string

	IsActive() bool
}

// IContext 上下文信息
type IContext interface {
	// Pixiver bind ctx
	Pixiver() IPixiver

	// ProxyUri transport use
	ProxyUri() string
}

// IApi api实例配置配置信息
type IApi interface {
	Method() string
	Url() string
}

// IClient http.request obj
type IClient interface {
	SetHeaders(...THeader) error
	SetCookies(string, ...TCookie) error
}

// IRequest request interface
type IRequest interface {
	Head(IApi, IContext, *TQuery, *TJson) (*http.Response, error)
	Get(IApi, IContext, *TQuery, *TJson) (*http.Response, error)
	Post(IApi, IContext, *TQuery, *TJson) (*http.Response, error)
	Put(IApi, IContext, *TQuery, *TJson) (*http.Response, error)
	Delete(IApi, IContext, *TQuery, *TJson) (*http.Response, error)
}

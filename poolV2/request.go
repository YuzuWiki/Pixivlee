package poolV2

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee"
)

type Request struct {
}

func (r Request) Head(api Pixivlee.IApi, ctx Pixivlee.TContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Get(api Pixivlee.IApi, ctx Pixivlee.TContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Post(api Pixivlee.IApi, ctx Pixivlee.TContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Put(api Pixivlee.IApi, ctx Pixivlee.TContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Delete(api Pixivlee.IApi, ctx Pixivlee.TContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

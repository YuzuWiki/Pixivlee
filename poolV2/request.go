package poolV2

import (
	"github.com/YuzuWiki/Pixivlee"
	"net/http"
)

type Request struct {
}

func (r Request) Head(api Pixivlee.IApi, ctx Pixivlee.IContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Get(api Pixivlee.IApi, ctx Pixivlee.IContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Post(api Pixivlee.IApi, ctx Pixivlee.IContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Put(api Pixivlee.IApi, ctx Pixivlee.IContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

func (r Request) Delete(api Pixivlee.IApi, ctx Pixivlee.IContext, query *Pixivlee.TQuery, json *Pixivlee.TJson) (*http.Response, error) {
	return nil, nil
}

package requests

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee/common"
)

func (r *requests) SetHeader(options ...common.HeaderOption) {
	if len(options) == 0 {
		return
	}

	r.BeforeHooks(
		func(req *http.Request) error {
			for _, option := range options {
				req.Header.Set(option.Key, option.Value)
			}
			return nil
		},
	)
}

package options

import (
	"time"

	"github.com/YuzuWiki/Pixivlee/kernel/middlewares"
	"github.com/YuzuWiki/Pixivlee/types"
)

func Initialize(container types.IKernel, o *HttpOption) error {
	if o == nil {
		return nil
	}
	api, host := formatUrl(o.BaseURl), formatUrl(o.Host)

	session := container.Session()

	session.SetBaseURL(api)

	// set debug
	session.SetDebug(o.IsDebug)

	// set timeout
	session.SetTimeout(time.Duration(o.Timeout) * time.Second)

	// set proxy
	session.SetProxy(o.ProxyUrl)

	// default request middleware
	session.OnBeforeRequest(middlewares.DefaultRequest(o.UserAgent, host, host))

	// default response middleware
	session.OnAfterResponse(middlewares.DefaultResponse())

	// config
	for _, fn := range o.OnBeforeRequest {
		session.OnBeforeRequest(fn)
	}

	for _, fn := range o.OnAfterResponse {
		session.OnAfterResponse(fn)
	}
	return nil
}

package Pixivlee

import (
	"fmt"
	"strings"

	"github.com/imroc/req/v3"

	fanbox "github.com/YuzuWiki/Pixivlee/fanbox"
	kernel "github.com/YuzuWiki/Pixivlee/kernel"
	pixiv "github.com/YuzuWiki/Pixivlee/pixiv"

	types "github.com/YuzuWiki/Pixivlee/types"
)

func formatUrl(host string) string {
	host = strings.TrimSpace(host)
	if len(host) == 0 {
		return ""
	}

	if strings.HasPrefix(host, "http") {
		return host
	}
	return fmt.Sprintf("https://%s", host)
}

func newContainer(api, host string, option kernel.Options) (types.IKernel, error) {
	api, host = formatUrl(api), formatUrl(host)

	container := kernel.NewKernel()

	container.SetBaseURL(api)

	if option.Http.IsDebug {
		container.EnableDebug()
	}

	if option.Http.Timeout > 0 {
		container.SetTimeOut(option.Http.Timeout)
	}

	if option.Http.ProxyUrl != "" {
		container.SetProxy(option.Http.ProxyUrl)
	}

	for _, fn := range option.Http.OnBeforeRequest {
		container.OnBeforeRequest(fn)
	}

	for _, fn := range option.Http.OnAfterResponse {
		container.OnAfterResponse(fn)
	}

	container.OnAfterResponse(func(client *req.Client, resp *req.Response) error {
		// Todo: do err status
		if HttpCode := resp.StatusCode; HttpCode != 200 {
			body, err := resp.ToBytes()
			if err != nil {
				return err
			}

			// eg. 400  {"error":"general_error"}
			return fmt.Errorf(fmt.Sprintf("%d  %s", HttpCode, string(body)))
		}
		return nil
	})

	if option.Http.UserAgent == "" {
		container.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			req.SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0").
				SetHeader("referer", host).
				SetHeader("origin", host)
			return nil
		})
	} else {
		container.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			req.SetHeader("User-Agent", strings.TrimSpace(option.Http.UserAgent)).
				SetHeader("referer", host).
				SetHeader("origin", host)
			return nil
		})
	}

	pixiver, err := types.NewPixiver(option.Pixiv.PhpSessId)
	if err != nil {
		return nil, err
	}
	container.SetPixiver(pixiver)
	return container, nil
}

func NewPixiv(option kernel.Options) (*pixiv.Client, error) {
	container, err := newContainer("www.pixiv.net", "www.pixiv.net", option)
	if err != nil {
		return nil, err
	}
	return pixiv.RegisterProvider(container), nil
}

func NewFanbox(option kernel.Options) (*fanbox.Client, error) {
	container, err := newContainer("api.fanbox.cc", "fanbox.cc", option)
	if err != nil {
		return nil, err
	}
	return fanbox.RegisterProvider(container), err
}

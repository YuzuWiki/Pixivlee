package Pixivlee

import (
	"github.com/YuzuWiki/Pixivlee/apis/fanbox"
	pixiv "github.com/YuzuWiki/Pixivlee/apis/pixiv"
	kernel "github.com/YuzuWiki/Pixivlee/kernel"
	types "github.com/YuzuWiki/Pixivlee/types"
)

func newContainer(api, host string, option kernel.Options) (types.IKernel, error) {
	container := kernel.NewKernel()

	// init session
	if err := option.Session.Initialize(container); err != nil {
		return nil, err
	}

	if err := option.Pixiv.Initialize(container); err != nil {
		return nil, err
	}
	return container, nil
}

func NewPixiv(option kernel.Options) (*pixiv.Client, error) {
	container, err := newContainer(BASE_URL_PIXIV, HOST_PIXIV, option)
	if err != nil {
		return nil, err
	}
	return pixiv.RegisterProvider(container), nil
}

func NewFanbox(option kernel.Options) (*fanbox.Client, error) {
	container, err := newContainer(BASE_URL_FANBOX, HOST_FANBOX, option)
	if err != nil {
		return nil, err
	}
	return fanbox.RegisterProvider(container), err
}

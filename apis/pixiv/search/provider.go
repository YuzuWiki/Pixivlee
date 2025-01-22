package search

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	Tags(JpName string) (*TagDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

package profile

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	All(pid types.TPid) (*AllDTO, error)
	Top(pid types.TPid) (*AllDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

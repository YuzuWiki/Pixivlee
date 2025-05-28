package information

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	Info(types.TPid) (*InfoDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

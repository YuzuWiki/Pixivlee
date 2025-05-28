package creator

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	Creator(username string) (*CreatorDTO, error)
	Get(pid types.TPid) (*InfoDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

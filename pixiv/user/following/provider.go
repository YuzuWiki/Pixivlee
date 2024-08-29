package following

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	List(pid types.TArtId, limit int, offset int) (*ListDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

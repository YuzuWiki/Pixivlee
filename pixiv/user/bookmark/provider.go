package bookmark

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	IllustShow(pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error)
	IllustHide(pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

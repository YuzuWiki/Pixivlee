package artwork

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	Illust(artId types.TArtId) (*ArtWorkDTO, error)
	Manga(artId types.TArtId) (*ArtWorkDTO, error)
	Novel(artId types.TArtId) (*ArtWorkDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

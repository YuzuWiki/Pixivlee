package ranking

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	ALl(mode string, page int, date string) (*RankDTO, error)
	Illust(mode string, page int, date string) (*RankDTO, error)
	Ugoira(mode string, page int, date string) (*RankDTO, error)
	Manga(mode string, page int, date string) (*RankDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

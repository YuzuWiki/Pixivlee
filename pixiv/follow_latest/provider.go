package follow_latest

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type IClient interface {
	Illust(page int) (*FollowLatestDTO, error)
	Manga(page int) (*FollowLatestDTO, error)
	Novel(page int) (*FollowLatestDTO, error)
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

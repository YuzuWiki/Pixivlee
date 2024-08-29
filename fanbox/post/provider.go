package post

import "github.com/YuzuWiki/Pixivlee/types"

type IClient interface {
	Info(PostId types.TPostId) (*PostDTO, error)
	List(creatorId string, limit int) TListIterator
}

func RegisterProvider(kernel types.IKernel) IClient {
	return &Client{kernel: kernel}
}

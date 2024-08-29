package fanbox

import (
	"github.com/YuzuWiki/Pixivlee/types"

	"github.com/YuzuWiki/Pixivlee/fanbox/creator"
	"github.com/YuzuWiki/Pixivlee/fanbox/post"
)

type Client struct {
	Creator creator.IClient
	Post    post.IClient
}

func RegisterProvider(kernel types.IKernel) *Client {
	return &Client{
		Creator: creator.RegisterProvider(kernel),
		Post:    post.RegisterProvider(kernel),
	}
}

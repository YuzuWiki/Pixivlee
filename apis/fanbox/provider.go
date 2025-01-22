package fanbox

import (
	"github.com/YuzuWiki/Pixivlee/apis/fanbox/creator"
	"github.com/YuzuWiki/Pixivlee/apis/fanbox/post"
	"github.com/YuzuWiki/Pixivlee/types"
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

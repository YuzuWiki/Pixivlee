package bookmark

import (
	"github.com/YuzuWiki/Pixivlee/types"

	"github.com/YuzuWiki/Pixivlee/pixiv/user/bookmark"
	"github.com/YuzuWiki/Pixivlee/pixiv/user/following"
	"github.com/YuzuWiki/Pixivlee/pixiv/user/information"
	"github.com/YuzuWiki/Pixivlee/pixiv/user/profile"
)

type Client struct {
	Information information.IClient
	Bookmark    bookmark.IClient
	Following   following.IClient
	Profile     profile.IClient
}

func RegisterProvider(kernel types.IKernel) *Client {
	return &Client{
		Information: information.RegisterProvider(kernel),
		Bookmark:    bookmark.RegisterProvider(kernel),
		Following:   following.RegisterProvider(kernel),
		Profile:     profile.RegisterProvider(kernel),
	}
}

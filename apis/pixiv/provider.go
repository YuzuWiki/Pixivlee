package bookmark

import (
	"github.com/YuzuWiki/Pixivlee/apis/pixiv/artwork"
	"github.com/YuzuWiki/Pixivlee/apis/pixiv/follow_latest"
	"github.com/YuzuWiki/Pixivlee/apis/pixiv/ranking"
	"github.com/YuzuWiki/Pixivlee/apis/pixiv/search"
	user "github.com/YuzuWiki/Pixivlee/apis/pixiv/user"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	User         *user.Client
	ArtWork      artwork.IClient
	FollowLatest follow_latest.IClient
	Ranking      ranking.IClient
	Search       search.IClient
}

func RegisterProvider(kernel types.IKernel) *Client {
	return &Client{
		User:         user.RegisterProvider(kernel),
		ArtWork:      artwork.RegisterProvider(kernel),
		FollowLatest: follow_latest.RegisterProvider(kernel),
		Ranking:      ranking.RegisterProvider(kernel),
		Search:       search.RegisterProvider(kernel),
	}
}

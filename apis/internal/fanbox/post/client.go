package post

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

// Info https://api.fanbox.cc/post.info?postId=5860717
func (c *Client) Info(PostId types.TPostId) (*PostDTO, error) {
	response := types.TFanboxResponse[PostDTO]{}

	r := c.kernel.NewRequests().SetResult(&response)

	if _, err := r.Get(fmt.Sprintf("/post.info?postId=%d", PostId)); err != nil {
		return nil, err
	}

	return response.Result()
}

// List "https://api.fanbox.cc/post.listCreator?creatorId=zhibujiloom&maxId=4160465&limit=100&maxPublishedDatetime=2025-01-01",
func (c *Client) List(creatorId string, limit int) TListIterator {

	return TListIterator{
		req:       c.kernel.NewRequests(),
		creatorId: creatorId,
		limit:     limit,
	}
}

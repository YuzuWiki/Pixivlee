package post

import (
	"fmt"
	"time"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

// Info https://api.fanbox.cc/post.info?postId=5860717
func (c *Client) Info(PostId types.TPostId) (*PostDTO, error) {
	response := types.TFanboxResponse[PostDTO]{}

	r := c.kernel.NewRequests().SetSuccessResult(&response)
	resp, err := r.Get(fmt.Sprintf("/post.info?postId=%d", PostId))
	if err != nil {
		return nil, err
	}
	resp.ErrorResult()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Request.URL)
	return response.Result()
}

// List "https://api.fanbox.cc/post.listCreator?creatorId=zhibujiloom&maxId=4160465&limit=100&maxPublishedDatetime=2025-01-01",
func (c *Client) List(creatorId string, limit int) TListIterator {
	r := c.kernel.NewRequests()
	return TListIterator{
		r:             r,
		creatorId:     creatorId,
		limit:         limit,
		publishedDate: time.Now().AddDate(1, 0, 0).Format(time.DateOnly),
	}
}

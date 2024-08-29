package follow_latest

import (
	"fmt"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

// followedLast (Last By Followed)
// page > 0 && data from https://www.pixiv.net/bookmark_new_illust.php
func followedLast(c *Client, mode string, page int) (*FollowLatestDTO, error) {
	data := types.TPixivResponse[FollowLatestDTO]{}

	r := c.kernel.NewRequests().
		AddQueryParam("pixiver", fmt.Sprint(page)).
		AddQueryParam("mode", mode).
		SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/follow_latest/%s", mode)); err != nil {
		return nil, err
	}

	return data.Result()
}

func (c *Client) Illust(page int) (*FollowLatestDTO, error) {
	return followedLast(c, "illust", page)
}

func (c *Client) Manga(page int) (*FollowLatestDTO, error) {
	return followedLast(c, "manga", page)
}

func (c *Client) Novel(page int) (*FollowLatestDTO, error) {
	return followedLast(c, "novel", page)
}

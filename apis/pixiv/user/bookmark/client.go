package bookmark

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func (c *Client) illust(rest string, pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	data := types.TPixivResponse[BookmarkDTO]{}

	r := c.kernel.NewRequests().
		SetQueryParam("tag", tag).
		SetQueryParam("limit", fmt.Sprint(limit)).
		SetQueryParam("offset", fmt.Sprint(offset)).
		SetQueryParam("rest", rest).
		SetResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/illusts/bookmarks", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

func (c *Client) IllustShow(pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	return c.illust("show", pid, tag, offset, limit)
}

func (c *Client) IllustHide(pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	return c.illust("hide", pid, tag, offset, limit)
}

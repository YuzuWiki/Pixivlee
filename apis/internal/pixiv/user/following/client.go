package following

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func (c *Client) List(pid types.TArtId, limit int, offset int) (*ListDTO, error) {
	data := types.TPixivResponse[ListDTO]{}

	r := c.kernel.NewRequests().
		SetPathParams(map[string]string{
			"offset": fmt.Sprint(offset),
			"limit":  fmt.Sprint(limit),
			"tag":    "",
			"rest":   "show",
		}).SetResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/following", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

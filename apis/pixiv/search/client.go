package search

import (
	"fmt"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func (c *Client) Tags(JpName string) (*TagDTO, error) {
	data := types.TPixivResponse[TagDTO]{}

	r := c.kernel.NewRequests().SetResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/search/tags/%s", JpName)); err != nil {
		return nil, err
	}

	return data.Result()
}

package information

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func (c *Client) Info(pid types.TPid) (*InfoDTO, error) {
	data := types.TPixivResponse[InfoDTO]{}

	r := c.kernel.NewRequests().SetSuccessResult(&data)

	r.SetPathParam("full", "1").SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

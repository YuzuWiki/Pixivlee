package profile

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func profile[T any](c *Client, mod string, pid types.TPid) (*T, error) {
	data := types.TPixivResponse[T]{}

	r := c.kernel.NewRequests().SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/profile/%s", pid, mod)); err != nil {
		return nil, err
	}
	return data.Result()
}

func (c *Client) All(pid types.TPid) (*AllDTO, error) {
	return profile[AllDTO](c, "all", pid)
}

func (c *Client) Top(pid types.TPid) (*AllDTO, error) {
	return profile[AllDTO](c, "top", pid)
}

package artwork

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func get[T any](c *Client, artType string, artIId types.TArtId) (*T, error) {
	data := types.TPixivResponse[T]{}

	r := c.kernel.NewRequests().SetResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/%s/%d", artType, artIId)); err != nil {
		return nil, err
	}
	return data.Result()
}

func (c *Client) Illust(artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](c, "illust", artId)
}

func (c *Client) Manga(artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](c, "manga", artId)
}

func (c *Client) Novel(artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](c, "novel", artId)
}

package ranking

import (
	"fmt"
	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func rank(c *Client, mode string, content string, page int, date string) (*RankDTO, error) {
	var data RankDTO

	r := c.kernel.NewRequests().
		SetSuccessResult(&data).
		AddQueryParam("page", fmt.Sprint(page)).
		AddQueryParam("format", "json")

	if len(mode) > 0 {
		r.AddQueryParam("mode", mode)
	}

	if len(content) > 0 {
		r.AddQueryParam("content", content)
	}

	if len(date) > 0 {
		r.AddQueryParam("date", date)
	}

	if _, err := r.Get(fmt.Sprintf("/ranking.php")); err != nil {
		return nil, err
	}

	return data.Result()
}

func (c *Client) ALl(mode string, page int, date string) (*RankDTO, error) {
	return rank(c, mode, "", page, date)
}

func (c *Client) Illust(mode string, page int, date string) (*RankDTO, error) {
	return rank(c, mode, "illust", page, date)
}

func (c *Client) Ugoira(mode string, page int, date string) (*RankDTO, error) {
	return rank(c, mode, "ugoira", page, date)
}

func (c *Client) Manga(mode string, page int, date string) (*RankDTO, error) {
	return rank(c, mode, "manga", page, date)
}

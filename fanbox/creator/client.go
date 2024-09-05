package creator

import (
	"fmt"
	"strings"

	"github.com/YuzuWiki/Pixivlee/types"
)

type Client struct {
	kernel types.IKernel
}

func (c *Client) Get(pid types.TPid) (*InfoDTO, error) {
	r := c.kernel.NewRequests()

	resp, err := r.Get(fmt.Sprintf("https://www.pixiv.net/fanbox/creator/%d", pid))
	if err != nil {
		return nil, err
	}

	// {CreatorName}.fanbox.cc
	CreatorUrl := resp.RawResponse().Request.URL.Host
	if !strings.HasSuffix(CreatorUrl, ".fanbox.cc") {
		return nil, fmt.Errorf("not found creator")
	}

	CreatorId, _ := strings.CutSuffix(CreatorUrl, ".fanbox.cc")
	return &InfoDTO{
		CreatorId:  CreatorId,
		CreatorUrl: CreatorUrl,
	}, nil
}

func (c *Client) Creator(username string) (*CreatorDTO, error) {
	username = strings.TrimSpace(username)
	if len(username) == 0 {
		return nil, fmt.Errorf("username is invalid, must not be empty")
	}
	url := fmt.Sprintf("https://%s.fanbox.cc", username)

	data := types.TPixivResponse[CreatorDTO]{}
	r := c.kernel.NewRequests().
		SetHeaders(map[string]string{"Referer": url, "Origin": url}).
		SetResult(&data)

	if _, err := r.Get(fmt.Sprintf("/creator.get?creatorId=%s", username)); err != nil {
		return nil, err
	}
	return data.Result()
}

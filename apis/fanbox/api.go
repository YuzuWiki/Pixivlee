package fanbox

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/types"
)

func CreatorUrl(pixiver Pixivlee.IPixiver, pid types.TArtId) (url string, err error) {
	r := Pixivlee.NewRequests(pixiver)

	resp, err := r.Get(fmt.Sprintf("/fanbox/creator/%d", pid))
	if err != nil {
		return "", err
	}
	return resp.Response.Request.URL.String(), nil
}

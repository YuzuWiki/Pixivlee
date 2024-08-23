package following

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/types"
)

func List(pixiver Pixivlee.IPixiver, pid types.TArtId, limit int, offset int) (*ListDTO, error) {
	data := apis.ResponseData[ListDTO]{}

	r := Pixivlee.NewRequests(pixiver).
		SetPathParams(map[string]string{
			"offset": fmt.Sprint(offset),
			"limit":  fmt.Sprint(limit),
			"tag":    "",
			"rest":   "show",
		}).SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/following", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

package account

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/types"
)

func Info(pixiver Pixivlee.IPixiver, pid types.TArtId) (*InfoDTO, error) {
	data := apis.ResponseData[InfoDTO]{}

	r := Pixivlee.NewRequests(pixiver).SetSuccessResult(&data)

	r.SetPathParam("full", "1").SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

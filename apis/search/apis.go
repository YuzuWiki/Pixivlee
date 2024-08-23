package search

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
)

func Tags(pixiver Pixivlee.IPixiver, JpName string) (*TagDTO, error) {
	data := apis.ResponseData[TagDTO]{}

	r := Pixivlee.NewRequests(pixiver).SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/search/tags/%s", JpName)); err != nil {
		return nil, err
	}

	return data.Result()
}

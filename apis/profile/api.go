package profile

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/types"
)

func profile[T any](mod string, pixiver Pixivlee.IPixiver, pid types.TPid) (*T, error) {
	data := apis.ResponseData[T]{}

	r := Pixivlee.NewRequests(pixiver).SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/profile/%s", pid, mod)); err != nil {
		return nil, err
	}
	return data.Result()
}

func All(pixiver Pixivlee.IPixiver, pid types.TPid) (*AllDTO, error) {
	return profile[AllDTO]("all", pixiver, pid)
}

func Top(pixiver Pixivlee.IPixiver, pid types.TPid) (*TopDTO, error) {
	return profile[TopDTO]("top", pixiver, pid)
}

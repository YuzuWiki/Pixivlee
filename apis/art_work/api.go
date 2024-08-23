package art_work

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/types"
)

func get[T any](pixiver Pixivlee.IPixiver, artType string, artIId types.TArtId) (*T, error) {
	data := apis.ResponseData[T]{}

	r := Pixivlee.NewRequests(pixiver).SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/%s/%d", artType, artIId)); err != nil {
		return nil, err
	}
	return data.Result()
}

func Illust(pixiver Pixivlee.IPixiver, artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](pixiver, "illust", artId)
}

func Manga(pixiver Pixivlee.IPixiver, artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](pixiver, "manga", artId)
}

func Novel(pixiver Pixivlee.IPixiver, artId types.TArtId) (*ArtWorkDTO, error) {
	return get[ArtWorkDTO](pixiver, "novel", artId)
}

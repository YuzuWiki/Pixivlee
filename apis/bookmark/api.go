package bookmark

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
	"github.com/YuzuWiki/Pixivlee/types"
)

func illust(pixiver Pixivlee.IPixiver, rest string, pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	data := apis.ResponseData[BookmarkDTO]{}

	r := Pixivlee.NewRequests(pixiver).
		AddQueryParam("tag", tag).
		AddQueryParam("limit", fmt.Sprint(limit)).
		AddQueryParams("offset", fmt.Sprint(offset)).
		AddQueryParams("rest", rest).
		SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/user/%d/illusts/bookmarks", pid)); err != nil {
		return nil, err
	}
	return data.Result()
}

func IllustShow(pixiver Pixivlee.IPixiver, pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	return illust(pixiver, "show", pid, tag, offset, limit)
}

func IllustHide(pixiver Pixivlee.IPixiver, pid types.TPid, tag string, offset, limit int) (*BookmarkDTO, error) {
	return illust(pixiver, "hide", pid, tag, offset, limit)
}

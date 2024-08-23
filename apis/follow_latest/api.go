package follow_latest

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/apis"
)

// followedLast (Last By Followed)
// page > 0 && data from https://www.pixiv.net/bookmark_new_illust.php
func followedLast(pixiver Pixivlee.IPixiver, mode string, page int) (*FollowLatestDTO, error) {
	data := apis.ResponseData[FollowLatestDTO]{}

	r := Pixivlee.NewRequests(pixiver).
		AddQueryParam("pixiver", fmt.Sprint(page)).
		AddQueryParam("mode", mode).
		SetSuccessResult(&data)

	if _, err := r.Get(fmt.Sprintf("/ajax/follow_latest/%s", mode)); err != nil {
		return nil, err
	}

	return data.Result()
}

func Illust(pixiver Pixivlee.IPixiver, page int) (*FollowLatestDTO, error) {
	return followedLast(pixiver, "illust", page)
}

func Manga(pixiver Pixivlee.IPixiver, page int) (*FollowLatestDTO, error) {
	return followedLast(pixiver, "manga", page)
}

func Novel(pixiver Pixivlee.IPixiver, page int) (*FollowLatestDTO, error) {
	return followedLast(pixiver, "novel", page)
}

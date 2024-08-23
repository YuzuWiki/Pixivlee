package ranking

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee"
)

func rank(pixiver Pixivlee.IPixiver, mode string, content string, page int, date string) (*RankDTO, error) {
	var data RankDTO

	r := Pixivlee.NewRequests(pixiver).
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

func ALl(pixiver Pixivlee.IPixiver, mode string, page int, date string) (*RankDTO, error) {
	return rank(pixiver, mode, "", page, date)
}

func Illust(pixiver Pixivlee.IPixiver, mode string, page int, date string) (*RankDTO, error) {
	return rank(pixiver, mode, "illust", page, date)
}

func Ugoira(pixiver Pixivlee.IPixiver, mode string, page int, date string) (*RankDTO, error) {
	return rank(pixiver, mode, "ugoira", page, date)
}

func Manga(pixiver Pixivlee.IPixiver, mode string, page int, date string) (*RankDTO, error) {
	return rank(pixiver, mode, "manga", page, date)
}

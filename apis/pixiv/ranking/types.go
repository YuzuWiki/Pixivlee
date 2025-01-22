package ranking

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/types"
)

type tRankItem struct {
	UserID   types.TPid `json:"user_id"`
	UserName string     `json:"user_name"`

	ID    types.TArtId `json:"illust_id"`
	Title string       `json:"title"`
	Url   string       `json:"url"`

	Tags       []string         `json:"tags"`
	PageCount  types.TCount     `json:"illust_page_count"`
	IllustType types.TArtType   `json:"illust_type"`
	UpdateDate types.TTimestamp `json:"illust_upload_timestamp"`

	Rank           types.TCount `json:"rank"`
	ViewCount      types.TCount `json:"view_count"`
	IsBookmarked   bool         `json:"is_bookmarked"`
	IsBookmarkable bool         `json:"bookmarkable"`
}

type RankDTO struct {
	error string

	Items   []tRankItem `json:"contents"`
	Mode    string      `json:"mode"`
	Content string      `json:"content"`

	HasPrev   bool `json:"prev"`
	RankTotal int  `json:"rank_total"`

	CurrPage int `json:"page"`
	NextPage int `json:"next"`

	PrevDate    string `json:"prev_date"`
	CurrDate    string `json:"date"`
	HasNextDate bool   `json:"next_date"`
}

func (d *RankDTO) Result() (*RankDTO, error) {
	if d.error != "" {
		return nil, fmt.Errorf(d.error)
	}

	return d, nil
}

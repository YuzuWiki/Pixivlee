package jsonObject

import (
	"strconv"
	"time"
)

type TTimestamp time.Time

func (t *TTimestamp) UnmarshalJSON(body []byte) error {
	if _t, err := strconv.ParseInt(string(body), 10, 64); err != nil {
		return err
	} else {
		*t = TTimestamp(time.Unix(_t, 0))
	}
	return nil
}

type rankItem struct {
	UserID   TPid   `json:"user_id"`
	UserName string `json:"user_name"`

	ID    TArtId `json:"illust_id"`
	Title string `json:"title"`
	Url   string `json:"url"`

	Tags       []string    `json:"tags"`
	PageCount  TCount      `json:"illust_page_count"`
	IllustType TIllustType `json:"illust_type"`
	UpdateDate TTimestamp  `json:"illust_upload_timestamp"`

	Rank         TCount `json:"rank"`
	ViewCount    TCount `json:"view_count"`
	IsBookmarked bool   `json:"is_bookmarked"`
	Bookmarkable bool   `json:"bookmarkable"`
}

type RankJson struct {
	error string

	Items   []rankItem `json:"contents"`
	Mode    string     `json:"mode"`
	Content string     `json:"content"`

	HasPrev   bool `json:"prev"`
	RankTotal int  `json:"rank_total"`

	CurrPage int `json:"page"`
	NextPage int `json:"next"`

	PrevDate    string `json:"prev_date"`
	CurrDate    string `json:"date"`
	HasNextDate bool   `json:"next_date"`
}

package art_work

import (
	"time"

	"github.com/YuzuWiki/Pixivlee/types"
)

type tTagItem struct {
	Jp string `json:"jp,omitempty"`
}

type tTag struct {
	Tags []tTagItem `json:"tags"`
}

/*
ArtWorkDTO

eg:
illustItem: https://www.pixiv.net/ajax/illust/90735220?lang=jp
mangaItem: https://www.pixiv.net/ajax/illust/28819260
novelItem: https://www.pixiv.net/ajax/novel/18132849
*/
type ArtWorkDTO struct {
	ArtId types.TArtId `json:"id,string"`
	Pid   types.TPid   `json:"userId,string"`

	// artwork into
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateDate  time.Time `json:"createDate"`
	UpdateDate  time.Time `json:"updateDate"`

	// artwork abstract
	PageCount     int64 `json:"pageCount"`
	LikeCount     int64 `json:"likeCount"`
	BookmarkCount int64 `json:"bookmarkCount"`
	ViewCount     int64 `json:"viewCount"`

	// tag abstract
	Tags tTag `json:"tags"`
}

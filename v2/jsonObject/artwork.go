package jsonObject

import (
	"time"
)

type abstractTag struct {
	Tags []struct {
		Jp string `json:"tag,omitempty"`
	} `json:"tags"`
}

/*
ArtWork
eg:

	illustItem: https://www.pixiv.net/ajax/illust/90735220?lang=jp
	mangaItem: https://www.pixiv.net/ajax/illust/28819260
	novelItem: https://www.pixiv.net/ajax/novel/18132849
*/
type ArtWork struct {
	ArtId TArtId `json:"id,string"`
	Pid   TPid   `json:"userId,string"`

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
	Tags abstractTag `json:"tags"`
}

type illustItem struct {
	ID              TArtId      `json:"id"`
	Title           string      `json:"title"`
	IllustType      TIllustType `json:"illustType"`
	URL             string      `json:"url"`
	Description     string      `json:"description"`
	Tags            []string    `json:"tags"`
	UserID          TPid        `json:"userId"`
	UserName        string      `json:"userName"`
	PageCount       int         `json:"pageCount"`
	CreateDate      time.Time   `json:"createDate"`
	UpdateDate      time.Time   `json:"updateDate"`
	ProfileImageURL string      `json:"profileImageUrl"`
}

type mangaItem struct {
	ID              TArtId      `json:"id"`
	Title           string      `json:"title"`
	IllustType      TIllustType `json:"illustType"`
	URL             string      `json:"url"`
	Description     string      `json:"description"`
	Tags            []string    `json:"tags"`
	UserID          TPid        `json:"userId"`
	UserName        string      `json:"userName"`
	PageCount       int         `json:"pageCount"`
	CreateDate      time.Time   `json:"createDate"`
	UpdateDate      time.Time   `json:"updateDate"`
	ProfileImageURL string      `json:"profileImageUrl"`
}

type novelItem struct {
	ID            TArtId    `json:"id"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	Tags          []string  `json:"tags"`
	UserID        TPid      `json:"userId"`
	UserName      string    `json:"userName"`
	TextCount     int       `json:"textCount"`
	Description   string    `json:"description"`
	BookmarkCount int       `json:"bookmarkCount"`
	CreateDate    time.Time `json:"createDate"`
	UpdateDate    time.Time `json:"updateDate"`
	SeriesID      string    `json:"seriesId"`
	SeriesTitle   string    `json:"seriesTitle"`
}

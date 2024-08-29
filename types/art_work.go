package types

import (
	"time"
)

type TIllustItem struct {
	ID              TArtId    `json:"id"`
	Title           string    `json:"title"`
	IllustType      TArtType  `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          TPid      `json:"userId"`
	UserName        string    `json:"userName"`
	PageCount       int       `json:"pageCount"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type TMangaItem struct {
	ID              TArtId    `json:"id"`
	Title           string    `json:"title"`
	IllustType      TArtType  `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          TPid      `json:"userId"`
	UserName        string    `json:"userName"`
	PageCount       int       `json:"pageCount"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type TNovelItem struct {
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

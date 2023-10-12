package jsonObject

import (
	"time"

	v2 "github.com/YuzuWiki/Pixivlee/v2"
)

type workItem struct {
	ID             v2.TArtId `json:"id,string"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Description    string    `json:"description"`
	UserID         v2.TPid   `json:"userId,string"`
	UserName       string    `json:"userName"`
	PageCount      int       `json:"pageCount"`
	IsBookmarkable bool      `json:"isBookmarkable"`
	Alt            string    `json:"alt"`
	CreateDate     time.Time `json:"createDate"`
	UpdateDate     time.Time `json:"updateDate"`
}

type Bookmark struct {
	Works []workItem `json:"works"`
	Total int        `json:"total"`
}

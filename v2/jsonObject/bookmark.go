package jsonObject

import (
	"time"
)

type workItem struct {
	ID             TArtId    `json:"id"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Description    string    `json:"description"`
	Pid            TPid      `json:"userId"`
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

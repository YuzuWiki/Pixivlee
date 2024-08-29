package bookmark

import (
	"time"

	"github.com/YuzuWiki/Pixivlee/types"
)

type TWorkItem struct {
	ID             types.TArtId `json:"id"`
	Title          string       `json:"title"`
	URL            string       `json:"url"`
	Description    string       `json:"description"`
	Pid            types.TPid   `json:"userId"`
	UserName       string       `json:"userName"`
	PageCount      types.TCount `json:"pageCount"`
	IsBookmarkable bool         `json:"isBookmarkable"`
	Alt            string       `json:"alt"`
	CreateDate     time.Time    `json:"createDate"`
	UpdateDate     time.Time    `json:"updateDate"`
}

type BookmarkDTO struct {
	Works []TWorkItem  `json:"works"`
	Total types.TCount `json:"total"`
}

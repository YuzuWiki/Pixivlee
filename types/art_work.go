package types

import (
	"encoding/json"
	"strconv"
	"time"
)

type TArtIds []TArtId

func (o *TArtIds) UnmarshalJSON(body []byte) error {
	var artDict map[string]any

	err := json.Unmarshal(body, &artDict)
	if err != nil {
		return err
	}

	artIds := make(TArtIds, 0)
	for artIdStr, _ := range artDict {
		artId, err := strconv.ParseInt(artIdStr, 10, 64)
		if err != nil {
			return err
		}

		artIds = append(artIds, TArtId(artId))
	}

	*o = artIds
	return nil
}

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

package dtos

import (
	"encoding/json"
	"time"
)

// ArtworkDTO Artwork
// illust: https://www.pixiv.net/ajax/illust/90735220?lang=jp
// manga: https://www.pixiv.net/ajax/illust/28819260
// novel: https://www.pixiv.net/ajax/novel/18132849
type ArtworkDTO struct {
	ArtId int64 `json:"id,string"`
	Pid   int64 `json:"userId,string"`

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

type abstractTag struct {
	Tags []struct {
		Jp string `json:"tag,omitempty"`
	} `json:"tags"`
}

// base data
type illustDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	IllustType      int       `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	PageCount       int       `json:"pageCount"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type illustMapDTO map[string]illustDTO

func (dto *illustMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]illustDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type mangaDTO struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	IllustType      int       `json:"illustType"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	Tags            []string  `json:"tags"`
	UserID          string    `json:"userId"`
	UserName        string    `json:"userName"`
	PageCount       int       `json:"pageCount"`
	CreateDate      time.Time `json:"createDate"`
	UpdateDate      time.Time `json:"updateDate"`
	ProfileImageURL string    `json:"profileImageUrl"`
}

type mangaMapDTO map[string]mangaDTO

func (dto *mangaMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]mangaDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

type novelDTO struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	URL           string    `json:"url"`
	Tags          []string  `json:"tags"`
	UserID        string    `json:"userId"`
	UserName      string    `json:"userName"`
	TextCount     int       `json:"textCount"`
	Description   string    `json:"description"`
	BookmarkCount int       `json:"bookmarkCount"`
	CreateDate    time.Time `json:"createDate"`
	UpdateDate    time.Time `json:"updateDate"`
	SeriesID      string    `json:"seriesId"`
	SeriesTitle   string    `json:"seriesTitle"`
}

type novelMapDTO map[string]novelDTO

func (dto *novelMapDTO) UnmarshalJSON(body []byte) error {
	// NOTE: 处理该字段无数据时为 数组
	if len(body) < 5 {
		return nil
	}

	data := map[string]novelDTO{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	*dto = data
	return nil
}

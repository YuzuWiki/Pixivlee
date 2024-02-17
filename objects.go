package Pixivlee

import (
	"time"
)

/*
============== account DTO ==============
*/
type userData struct {
	Name         string `json:"name"`
	PrivacyLevel string `json:"privacyLevel"`
}

type User struct {
	UserID     TPid     `json:"userId,string"`
	Name       string   `json:"name"`
	Avatar     string   `json:"imageBig"`
	IsFollowed bool     `json:"isFollowed"`
	Following  int32    `json:"following"`
	Region     userData `json:"region"`
	Gender     userData `json:"gender"`
	BirthDay   userData `json:"birthDay"`
	Job        userData `json:"job"`
}

/*
============== artwork DTO ==============
*/
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

/*
============== bookmark DTO ==============
*/
type workItem struct {
	ID             TArtId    `json:"id"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Description    string    `json:"description"`
	Pid            TPid      `json:"userId"`
	UserName       string    `json:"userName"`
	PageCount      TCount    `json:"pageCount"`
	IsBookmarkable bool      `json:"isBookmarkable"`
	Alt            string    `json:"alt"`
	CreateDate     time.Time `json:"createDate"`
	UpdateDate     time.Time `json:"updateDate"`
}

type Bookmark struct {
	Works []workItem `json:"works"`
	Total TCount     `json:"total"`
}

/*
	============== following DTO ==============
*/
// Following  user
type Following struct {
	Total TCount       `json:"total"`
	Users []followUser `json:"users"`
}

type followUser struct {
	UserID        TPid         `json:"userId,string"`
	Name          string       `json:"userName"`
	Avatar        string       `json:"profileImageUrl"`
	UserComment   string       `json:"userComment"`
	IsFollowing   bool         `json:"following"`
	IsFollowed    bool         `json:"followed"`
	IsBlocking    bool         `json:"isBlocking"`
	IsMypixiv     bool         `json:"isMypixiv"`
	Illusts       []illustItem `json:"illusts"`
	AcceptRequest bool         `json:"acceptRequest"`
}

type thumbnailItem struct {
	Illust []illustItem `json:"illust"`
	Novel  []novelItem  `json:"novel"`
}

type page struct {
	Ids  []TArtId      `json:"ids"`
	Tags []interface{} `json:"tags"`
}

// FollowLatestDTO  follow_latest
type FollowLatestDTO struct {
	Page           page                              `json:"page"`
	TagTranslation jsonMap[string, multilingualItem] `json:"tagTranslation"`
	Thumbnails     thumbnailItem                     `json:"thumbnails"`
}

/*
============== profile DTO ==============
*/
type extra struct {
	Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Canonical   string `json:"canonical"`
		Ogp         struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Type        string `json:"type"`
		} `json:"ogp"`
		Twitter struct {
			Description string `json:"description"`
			Image       string `json:"image"`
			Title       string `json:"title"`
			Card        string `json:"card"`
		} `json:"twitter"`
		DescriptionHeader string `json:"descriptionHeader"`
	} `json:"meta"`
}

// ProfileTop return user's  profile (top)
type ProfileTop struct {
	Illusts   jsonMap[string, illustItem] `json:"illusts"`
	Manga     jsonMap[string, mangaItem]  `json:"manga"`
	Novels    jsonMap[string, novelItem]  `json:"novels"`
	ExtraData extra                       `json:"extra_data"`
}

// ProfileAll return user's  profile (all)
type ProfileAll struct {
	Illusts artWorkIds `json:"illusts"`
	Manga   artWorkIds `json:"mangas"`
	Novel   artWorkIds `json:"novels"`
}

/*
	============== ranking DTO ==============
*/

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

	Rank           TCount `json:"rank"`
	ViewCount      TCount `json:"view_count"`
	IsBookmarked   bool   `json:"is_bookmarked"`
	IsBookmarkable bool   `json:"bookmarkable"`
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

/*
	============== tag DTO ==============
*/

type multilingualItem struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

type tagDigest struct {
	Id           TArtId   `json:"id,string"`
	Abstract     string   `json:"abstract"`
	Image        string   `json:"image"`
	ParentTag    string   `json:"parentTag"`
	SiblingsTags []string `json:"siblingsTags"`
	ChildrenTags []string `json:"childrenTags"`
}

// Tag  https://www.pixiv.net/ajax/search/tags/%E4%BA%8C%E6%AC%A1%E5%89%B5%E4%BD%9C?lang=zh
type Tag struct {
	Jp          string                            `json:"tag"`
	Digest      tagDigest                         `json:"pixpedia"`
	Translation jsonMap[string, multilingualItem] `json:"tagTranslation"`
}

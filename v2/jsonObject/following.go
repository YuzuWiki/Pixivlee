package jsonObject

import v2 "github.com/YuzuWiki/Pixivlee/v2"

// Following  user
type Following struct {
	Total int          `json:"total"`
	Users []followUser `json:"users"`
}

type followUser struct {
	UserID        v2.TPid      `json:"userId,string"`
	Name          string       `json:"userName"`
	Avatar        string       `json:"profileImageUrl"`
	UserComment   string       `json:"userComment"`
	Following     bool         `json:"following"`
	Followed      bool         `json:"followed"`
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
	Ids  []v2.TArtId   `json:"ids"`
	Tags []interface{} `json:"tags"`
}

// FollowLatestDTO  follow_latest
type FollowLatestDTO struct {
	Page           page                              `json:"page"`
	TagTranslation jsonMap[string, multilingualItem] `json:"tagTranslation"`
	Thumbnails     thumbnailItem                     `json:"thumbnails"`
}
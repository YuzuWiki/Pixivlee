package dtos

// FollowingDTO  following user
type FollowingDTO struct {
	Total int             `json:"total"`
	Users []followUserDTO `json:"users"`
}

type followUserDTO struct {
	UserID        int64       `json:"userId,string"`
	Name          string      `json:"userName"`
	Avatar        string      `json:"profileImageUrl"`
	UserComment   string      `json:"userComment"`
	Following     bool        `json:"following"`
	Followed      bool        `json:"followed"`
	IsBlocking    bool        `json:"isBlocking"`
	IsMypixiv     bool        `json:"isMypixiv"`
	Illusts       []illustDTO `json:"illusts"`
	AcceptRequest bool        `json:"acceptRequest"`
}

// FollowLatestDTO  follow_latest
type FollowLatestDTO struct {
	Page           pageDTO      `json:"page"`
	TagTranslation tagTranslDTO `json:"tagTranslation"`
	Thumbnails     thumbnailDTO `json:"thumbnails"`
}

type pageDTO struct {
	Ids  []int64       `json:"ids"`
	Tags []interface{} `json:"tags"`
}

type thumbnailDTO struct {
	Illust []illustDTO `json:"illust"`
	Novel  []novelDTO  `json:"novel"`
}

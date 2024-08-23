package follow_latest

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type tPage struct {
	Ids  []types.TArtId `json:"ids"`
	Tags []interface{}  `json:"tags"`
}

type tMultilingualItem struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

type tThumbnailItem struct {
	Illust []types.TIllustItem `json:"illust"`
	Novel  []types.TNovelItem  `json:"novel"`
}

// FollowLatestDTO  follow_latest
type FollowLatestDTO struct {
	Page           tPage                                      `json:"page"`
	TagTranslation types.TJsonDict[string, tMultilingualItem] `json:"tagTranslation"`
	Thumbnails     tThumbnailItem                             `json:"thumbnails"`
}

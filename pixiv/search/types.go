package search

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type tTagDigest struct {
	Id           types.TArtId `json:"id,string"`
	Abstract     string       `json:"abstract"`
	Image        string       `json:"image"`
	ParentTag    string       `json:"parentTag"`
	SiblingsTags []string     `json:"siblingsTags"`
	ChildrenTags []string     `json:"childrenTags"`
}

type tTranslationItem struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

// TagDTO  https://www.pixiv.net/ajax/search/tags/%E4%BA%8C%E6%AC%A1%E5%89%B5%E4%BD%9C?lang=zh
type TagDTO struct {
	Jp          string                                    `json:"tag"`
	Digest      tTagDigest                                `json:"pixpedia"`
	Translation types.TJsonDict[string, tTranslationItem] `json:"tagTranslation"`
}

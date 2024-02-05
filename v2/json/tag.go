package json

import v2 "github.com/YuzuWiki/Pixivlee/v2"

type multilingualItem struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

type tagDigest struct {
	Id           v2.TArtId `json:"id,string"`
	Abstract     string    `json:"abstract"`
	Image        string    `json:"image"`
	ParentTag    string    `json:"parentTag"`
	SiblingsTags []string  `json:"siblingsTags"`
	ChildrenTags []string  `json:"childrenTags"`
}

// Tag  https://www.pixiv.net/ajax/search/tags/%E4%BA%8C%E6%AC%A1%E5%89%B5%E4%BD%9C?lang=zh
type Tag struct {
	Jp          string                            `json:"tag"`
	Digest      tagDigest                         `json:"pixpedia"`
	Translation jsonMap[string, multilingualItem] `json:"tagTranslation"`
}

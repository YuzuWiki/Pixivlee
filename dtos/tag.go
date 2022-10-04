package dtos

import (
	"encoding/json"
)

// TagDTO  https://www.pixiv.net/ajax/search/tags/%E4%BA%8C%E6%AC%A1%E5%89%B5%E4%BD%9C?lang=zh
type TagDTO struct {
	Jp          string       `json:"tag"`
	Digest      tagDigestDTO `json:"pixpedia"`
	Translation tagTranslDTO `json:"tagTranslation"`
}

type tagDigestDTO struct {
	Id           int64    `json:"id,string"`
	Abstract     string   `json:"abstract"`
	Image        string   `json:"image"`
	ParentTag    string   `json:"parentTag"`
	SiblingsTags []string `json:"siblingsTags"`
	ChildrenTags []string `json:"childrenTags"`
}

type multilingualDTO struct {
	En     string `json:"en"`
	Ko     string `json:"ko"`
	Zh     string `json:"zh"`
	Romaji string `json:"romaji"`
}

type tagTranslDTO map[string]multilingualDTO

func (dto *tagTranslDTO) UnmarshalJSON(body []byte) error {
	var data map[string]multilingualDTO
	if err := json.Unmarshal(body, &data); err == nil {
		*dto = data
	}
	return nil
}

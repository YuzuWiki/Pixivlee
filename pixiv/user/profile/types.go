package profile

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type AllDTO struct {
	Illusts types.TArtIds `json:"Illusts"`
	Mangas  types.TArtIds `json:"mangas"`
	Novels  types.TArtIds `json:"novels"`
}

type tOgp struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Type        string `json:"type"`
}

type tTwitter struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Card        string `json:"card"`
}

type tMeta struct {
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Canonical         string   `json:"canonical"`
	Ogp               tOgp     `json:"ogp"`
	Twitter           tTwitter `json:"twitter"`
	DescriptionHeader string   `json:"descriptionHeader"`
}

type extraDto struct {
	Meta tMeta `json:"meta"`
}

type TopDTO struct {
	Illusts   types.TJsonDict[string, types.TIllustItem] `json:"illusts"`
	Manga     types.TJsonDict[string, types.TMangaItem]  `json:"manga"`
	Novels    types.TJsonDict[string, types.TNovelItem]  `json:"novels"`
	ExtraData extraDto                                   `json:"extra_data"`
}

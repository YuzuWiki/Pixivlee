package creator

import "github.com/YuzuWiki/Pixivlee/types"

type TUser struct {
	UserId  types.TPid `json:"userId"`
	Name    string     `json:"name"`
	IconUrl string     `json:"iconUrl"`
}

type CreatorDTO struct {
	User               TUser    `json:"user"`
	CreatorId          string   `json:"creatorId"`
	Description        string   `json:"description"`
	HasAdultContent    bool     `json:"hasAdultContent"`
	CoverImageUrl      string   `json:"coverImageUrl"`
	ProfileLinks       []string `json:"profileLinks"`
	IsFollowed         bool     `json:"isFollowed"`
	IsSupported        bool     `json:"isSupported"`
	IsAcceptingRequest bool     `json:"isAcceptingRequest"`
	HasBoothShop       bool     `json:"hasBoothShop"`
	HasPublishedPost   bool     `json:"hasPublishedPost"`
}

type InfoDTO struct {
	CreatorId  string `json:"creatorId"`
	CreatorUrl string `json:"CreatorUrl"`
}

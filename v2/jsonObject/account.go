package jsonObject

import v2 "github.com/YuzuWiki/Pixivlee/v2"

type info struct {
	Name         string `json:"name"`
	PrivacyLevel string `json:"privacyLevel"`
}

type User struct {
	UserID     v2.TPid `json:"userId,string"`
	Name       string  `json:"name"`
	Avatar     string  `json:"imageBig"`
	IsFollowed bool    `json:"isFollowed"`
	Following  int32   `json:"following"`
	Region     info    `json:"region"`
	Gender     info    `json:"gender"`
	BirthDay   info    `json:"birthDay"`
	Job        info    `json:"job"`
}

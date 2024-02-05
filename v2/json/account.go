package json

import v2 "github.com/YuzuWiki/Pixivlee/v2"

type userData struct {
	Name         string `json:"name"`
	PrivacyLevel string `json:"privacyLevel"`
}

type User struct {
	UserID     v2.TPid  `json:"userId,string"`
	Name       string   `json:"name"`
	Avatar     string   `json:"imageBig"`
	IsFollowed bool     `json:"isFollowed"`
	Following  int32    `json:"following"`
	Region     userData `json:"region"`
	Gender     userData `json:"gender"`
	BirthDay   userData `json:"birthDay"`
	Job        userData `json:"job"`
}

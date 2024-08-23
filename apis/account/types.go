package account

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type tUserData struct {
	Name         string `json:"name"`
	PrivacyLevel string `json:"privacyLevel"`
}

type InfoDTO struct {
	UserID     types.TPid `json:"userId,string"`
	Name       string     `json:"name"`
	Avatar     string     `json:"imageBig"`
	IsFollowed bool       `json:"isFollowed"`
	Following  int32      `json:"following"`
	Region     tUserData  `json:"region"`
	Gender     tUserData  `json:"gender"`
	BirthDay   tUserData  `json:"birthDay"`
	Job        tUserData  `json:"job"`
}

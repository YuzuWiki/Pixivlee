package dtos

type UserInfoDTO struct {
	UserID     int64       `json:"userId,string"`
	Name       string      `json:"name"`
	Avatar     string      `json:"imageBig"`
	IsFollowed bool        `json:"isFollowed"`
	Following  int32       `json:"following"`
	Region     userDataDTO `json:"region"`
	Gender     userDataDTO `json:"gender"`
	BirthDay   userDataDTO `json:"birthDay"`
	Job        userDataDTO `json:"job"`
}

type userDataDTO struct {
	Name         string `json:"name"`
	PrivacyLevel string `json:"privacyLevel"`
}

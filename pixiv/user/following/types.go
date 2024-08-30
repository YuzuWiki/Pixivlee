package following

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type tFollowUser struct {
	UserID        types.TPid          `json:"userId,string"`
	Name          string              `json:"userName"`
	Avatar        string              `json:"profileImageUrl"`
	UserComment   string              `json:"userComment"`
	IsFollowing   bool                `json:"following"`
	IsFollowed    bool                `json:"followed"`
	IsBlocking    bool                `json:"isBlocking"`
	IsMypixiv     bool                `json:"isMypixiv"`
	Illusts       []types.TIllustItem `json:"illusts"`
	AcceptRequest bool                `json:"acceptRequest"`
}

type ListDTO struct {
	Total types.TCount  `json:"total"`
	Users []tFollowUser `json:"users"`
}

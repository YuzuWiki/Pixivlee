package post

import (
	"fmt"

	"github.com/YuzuWiki/Pixivlee/kernel/request"
	"github.com/YuzuWiki/Pixivlee/types"
)

type TUser struct {
	UserId  types.TPid `json:"userId"`
	Name    string     `json:"name"`
	IconUrl string     `json:"iconUrl"`
}

type TCover struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type TSimplePost struct {
	Id                types.TPostId   `json:"id"`
	Title             string          `json:"title"`
	PublishedDatetime types.TTimeDate `json:"publishedDatetime"`
}

type PostDTO struct {
	Id                types.TPostId   `json:"id"`
	Title             string          `json:"title"`
	CreatorId         string          `json:"creatorId"`
	Type              string          `json:"type"`
	CoverImageUrl     string          `json:"coverImageUrl"`
	ImageForShare     string          `json:"imageForShare"`
	Excerpt           string          `json:"excerpt"`
	FeeRequired       int             `json:"feeRequired"`
	LikeCount         int             `json:"likeCount"`
	CommentCount      int             `json:"commentCount"`
	IsLiked           bool            `json:"isLiked"`
	IsRestricted      bool            `json:"isRestricted"`
	IsPinned          bool            `json:"isPinned"`
	HasAdultContent   bool            `json:"hasAdultContent"`
	PublishedDatetime types.TTimeDate `json:"publishedDatetime"`
	UpdatedDatetime   types.TTimeDate `json:"updatedDatetime"`
	Tags              []string        `json:"tags"` //  fixme
	User              TUser           `json:"user"`
	NextPost          TSimplePost     `json:"nextPost"`
	PrevPost          TSimplePost     `json:"prevPost"`
	Body              any             `json:"body"` //  fixme
	//commentList
}

type ListPostDTO struct {
	Id                types.TPostId   `json:"id"`
	Title             string          `json:"title"`
	CreatorId         string          `json:"creatorId"`
	FeeRequired       int             `json:"feeRequired"`
	LikeCount         int             `json:"likeCount"`
	CommentCount      int             `json:"commentCount"`
	Excerpt           string          `json:"excerpt"`
	IsLiked           bool            `json:"isLiked"`
	IsRestricted      bool            `json:"isRestricted"`
	IsPinned          bool            `json:"isPinned"`
	HasAdultContent   bool            `json:"hasAdultContent"`
	PublishedDatetime types.TTimeDate `json:"publishedDatetime"`
	UpdatedDatetime   types.TTimeDate `json:"updatedDatetime"`
	Tags              []string        `json:"tags"` //  fixme
	Cover             *TCover         `json:"cover,omitempty"`
	User              TUser           `json:"user"`
}

type ListDTO []ListPostDTO

type TListIterator struct {
	req request.IRequest

	creatorId     string
	limit         int
	publishedDate string
}

func (i *TListIterator) Next() func(yield func(idx int, items *ListDTO) bool) {
	return func(yield func(idx int, item *ListDTO) bool) {
		var (
			maxId types.TPostId = 0
			page                = 0
		)

		url := fmt.Sprintf("/post.listCreator?creatorId=%s&limit=%d", i.creatorId, i.limit)
		for {
			page += 1
			response := types.TFanboxResponse[ListDTO]{}

			if _, err := i.req.SetResult(&response).Get(url); err != nil {
				break
			}

			data, err := response.Result()
			if err != nil {
				break
			}

			length := len(*data)
			if length == 0 {
				break
			}
			yield(page, data)

			maxId = (*data)[length-1].Id
			url = fmt.Sprintf("/post.listCreator?creatorId=%s&limit=%d&maxPublishedDatetime=%s&maxId=%d", i.creatorId, i.limit, i.publishedDate, maxId)
		}
		return
	}
}

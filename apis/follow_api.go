package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func followArtWork(ctx common.IContext, mode string, page int) (body *dtos.FollowLatestDTO, err error) {
	var (
		query *common.Query
		c     common.IClient
	)

	if query, err = common.NewQuery(map[string]interface{}{"p": page, "mode": mode, "lang": "jp"}); err != nil {
		return
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := common.Json(c.Get, common.Path("/ajax/follow_latest", mode), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func FollowIllusts(ctx common.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Illust, page)
}

func FollowNovel(ctx common.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Novel, page)
}

func FollowManga(ctx common.IContext, page int) (*dtos.FollowLatestDTO, error) {
	return followArtWork(ctx, Manga, page)
}

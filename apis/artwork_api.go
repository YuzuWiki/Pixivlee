package apis

import (
	"encoding/json"
	"strconv"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func GetAccountPid(ctx common.IContext) (int64, error) {
	c, err := Pixivlee.Pool().Get(ctx.PhpSessID())
	if err != nil {
		return 0, err
	}

	header, err := common.Header(c.Get, "https://"+common.PixivHost, nil, nil)
	if err != nil {
		return 0, err
	}

	if pid := header.Get("x-userid"); len(pid) > 0 {
		return strconv.ParseInt(pid, 10, 64)
	}
	return 0, nil
}

func getArtWork(ctx common.IContext, artType string, artId int64) (_ *dtos.ArtworkDTO, err error) {
	var (
		query *common.Query
		c     common.IClient
		body  dtos.ArtworkDTO
	)

	if query, err = common.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := common.Json(c.Get, common.Path("/ajax", artType, artId), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	return &body, nil
}

func GetIllusts(ctx common.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Illust, artId)
}

func GetMangas(ctx common.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Manga, artId)
}

func GetNovels(ctx common.IContext, artId int64) (*dtos.ArtworkDTO, error) {
	return getArtWork(ctx, Novel, artId)
}

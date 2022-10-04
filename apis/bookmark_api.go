package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func bookmark(ctx common.IContext, rest string, uid int64, tag string, offset int, limit int) (body *dtos.BookmarkDTO, err error) {
	var (
		query *common.Query
		c     common.IClient
	)
	if query, err = common.NewQuery(map[string]interface{}{
		"tag":    tag,
		"limit":  limit,
		"offset": offset,
		"rest":   rest,
		"lang":   "jp",
	}); err != nil {
		return
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := common.Json(c.Get, common.Path("/ajax/user", uid, "/illusts/bookmarks"), query, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, body); err != nil {
		return
	}
	return body, err
}

func BookmarkShow(ctx common.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Show, uid, tag, offset, limit)
}

func BookmarkHide(ctx common.IContext, uid int64, tag string, offset int, limit int) (*dtos.BookmarkDTO, error) {
	return bookmark(ctx, Hide, uid, tag, offset, limit)
}

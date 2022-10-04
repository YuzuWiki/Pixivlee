package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func GetFollowing(ctx common.IContext, pid int64, limit int, offset int) (*dtos.FollowingDTO, error) {
	var (
		query *common.Query
		c     common.IClient
		err   error
	)

	if query, err = common.NewQuery(map[string]interface{}{
		"offset": offset,
		"limit":  limit,
		"rest":   Show,
		"tag":    "",
		"lang":   "jp",
	}); err != nil {
		return nil, err
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := common.Json(c.Get, common.Path("/ajax/user", pid, "/following"), query, nil)
	if err != nil {
		return nil, err
	}

	body := &dtos.FollowingDTO{}
	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

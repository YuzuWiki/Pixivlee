package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func GetTag(ctx common.IContext, jP string) (*dtos.TagDTO, error) {
	var (
		query *common.Query
		c     common.IClient
		err   error
		tag   dtos.TagDTO
	)

	if query, err = common.NewQuery(map[string]interface{}{"lang": "jp"}); err != nil {
		return nil, err
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := common.Json(c.Get, common.Path("/ajax/search/tags", jP), query, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &tag); err != nil {
		return nil, err
	}
	return &tag, nil
}

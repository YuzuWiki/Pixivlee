package apis

import (
	"encoding/json"

	"github.com/YuzuWiki/Pixivlee"
	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/dtos"
)

func GetAccountInfo(ctx common.IContext, pid int64) (body *dtos.UserInfoDTO, err error) {
	var (
		query *common.Query
		c     common.IClient
	)
	if query, err = common.NewQuery(map[string]interface{}{"lang": "jp", "full": 1}); err != nil {
		return
	}

	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := common.Json(c.Get, common.Path("/ajax/user", pid), query, nil)
	if err != nil {
		return nil, err
	}

	body = &dtos.UserInfoDTO{}
	if err = json.Unmarshal(data, body); err != nil {
		return nil, err
	}
	return body, nil
}

func GetProfileAll(ctx common.IContext, pid int64) (_ *dtos.AllProfileDTO, err error) {
	var (
		c    common.IClient
		body dtos.AllProfileDTO
	)
	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return nil, err
	}

	data, err := common.Json(c.Get, common.Path("/ajax/user/", pid, "/profile", All), nil, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &body); err != nil {
		return nil, err
	}

	return &body, nil
}

func GetProfileTop(ctx common.IContext, pid int64) (_ *dtos.TopProfileDTO, err error) {
	var (
		c    common.IClient
		body dtos.TopProfileDTO
	)
	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return
	}

	data, err := common.Json(c.Get, common.Path("/ajax/user/", pid, "/profile", Top), nil, nil)
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &body); err != nil {
		return
	}

	return &body, nil
}

func GetFanboxUlr(ctx common.IContext, pid int64) (u string, err error) {
	var c common.IClient
	if c, err = Pixivlee.Pool().Get(ctx.PhpSessID()); err != nil {
		return "", err
	}

	header, err := common.Header(c.Get, common.Path("/fanbox/creator/", pid), nil, nil)
	if err != nil {
		return "", err
	}
	return header.Get("location"), nil
}

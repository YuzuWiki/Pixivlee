package Pixivlee

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/imroc/req/v3"

	"github.com/YuzuWiki/Pixivlee/types"
)

type IPixiver interface {
	Pid() types.TPid
	SessionID() string
}

type tPixiverV2 struct {
	pid       types.TPid
	sessionId string
}

func (p *tPixiverV2) Pid() types.TPid {
	return p.pid
}

func (p *tPixiverV2) SessionID() string {
	return p.sessionId
}

func NewPixiver(sessionId string) (IPixiver, error) {
	sessionId = strings.TrimSpace(sessionId)
	if len(sessionId) == 0 {
		return nil, fmt.Errorf("sessionid is invalid, must not be empty!")
	}

	pixiver := &tPixiverV2{
		sessionId: sessionId,
	}

	arr := strings.SplitN(strings.TrimSpace(sessionId), "_", 2)
	if len(arr) != 2 {
		return nil, fmt.Errorf("sessionid is invalid, format error!")
	}

	pid, err := strconv.ParseUint(arr[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("sessionid is invalid, (pid) %s", err.Error())
	}
	pixiver.pid = types.TPid(pid)

	return pixiver, nil
}

func NewRequests(p IPixiver) *req.Request {
	r := requests.NewRequest()

	// set cookie
	r.SetCookies(
		&http.Cookie{
			Name:   PHPSESSID,
			Value:  p.SessionID(),
			Path:   "/",
			Domain: PIXIV_DOMAIN,
		})

	// default params
	r.AddQueryParam("lang", "jp")

	return r
}

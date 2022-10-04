package Pixivlee

import (
	"net/http"
	"strings"
	"sync"

	"github.com/YuzuWiki/Pixivlee/common"
	"github.com/YuzuWiki/Pixivlee/requests"
)

func newClient(sessionId string) (common.IClient, error) {
	c := requests.NewRequest()

	// set default header
	c.SetHeader(
		common.HeaderOption{Key: "User-Agent", Value: common.UserAgent},
		common.HeaderOption{Key: "referer", Value: "https://" + common.PixivHost},
	)

	// set cookie
	if len(sessionId) > 0 {
		if err := c.SetCookies(
			common.PixivHost,
			&http.Cookie{
				Name:   common.Phpsessid,
				Value:  sessionId,
				Path:   "/",
				Domain: common.PixivDomain,
			},
		); err != nil {
			return nil, err
		}
	}
	return c, nil
}

type pixivPool struct {
	p map[string]common.IClient

	m sync.Mutex
}

func (s *pixivPool) Get(sessionId string) (common.IClient, error) {
	var (
		isOk bool
		c    common.IClient
		err  error
	)

	sessionId = strings.TrimSpace(sessionId)
	if c, isOk = s.p[sessionId]; isOk {
		return c, nil
	}

	s.m.Lock()
	defer s.m.Unlock()

	if c, isOk = s.p[sessionId]; isOk {
		return c, nil
	}

	if c, err = newClient(sessionId); err != nil {
		return nil, err
	}

	s.p[sessionId] = c
	return c, nil
}

func (s *pixivPool) Del(sessionId string) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.p, sessionId)
}

func newPool() common.IPool {
	return &pixivPool{p: map[string]common.IClient{}}
}

func Pool() common.IPool {
	return pool
}

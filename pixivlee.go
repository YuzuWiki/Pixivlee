package Pixivlee

import (
	"fmt"
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
	m sync.Mutex

	// manage  session
	sessions map[string]common.IClient
}

func (s *pixivPool) Add(sessionId string) (c common.IClient, err error) {
	sessionId = strings.TrimSpace(sessionId)

	s.m.Lock()
	defer s.m.Unlock()

	if _, isOk := s.sessions[sessionId]; isOk {
		return nil, fmt.Errorf("the pixiver already exists")
	}

	if c, err = newClient(sessionId); err != nil {
		return nil, err
	}

	s.sessions[sessionId] = c
	return c, nil
}

func (s *pixivPool) Get(sessionId string) (common.IClient, error) {
	sessionId = strings.TrimSpace(sessionId)
	if c, isOk := s.sessions[sessionId]; isOk {
		return c, nil
	}
	return nil, fmt.Errorf("non-existent pixiver")

}

func (s *pixivPool) Del(sessionId string) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.sessions, sessionId)
}

func newPool() common.IPool {
	return &pixivPool{
		sessions: map[string]common.IClient{}}
}

func Pool() common.IPool {
	return pool
}

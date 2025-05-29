package kernel

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/YuzuWiki/Pixivlee/types"
)

type tBaseAccount struct {
	pid types.TPid

	sessionId string
}

func (p *tBaseAccount) Pid() types.TPid {
	return p.pid
}

func (p *tBaseAccount) SessionID() string {
	return p.sessionId
}

func (p *tBaseAccount) SetSessionID(sessionId string) error {
	sessionId = strings.TrimSpace(sessionId)
	if sessionId == "" {
		return fmt.Errorf("invalid sessionid")
	}

	arr := strings.SplitN(strings.TrimSpace(sessionId), "_", 2)
	switch len(arr) {
	case 1:
		p.sessionId = arr[0]
	case 2:
		pid, err := strconv.ParseUint(arr[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid session' pid")
		}
		p.pid = types.TPid(pid)

		p.sessionId = arr[1]
	default:
		return fmt.Errorf("unknown error")
	}
	return nil
}

type TPixiver struct {
	tBaseAccount
}

func (p *TPixiver) Cookies() []*http.Cookie {
	return []*http.Cookie{
		{
			Name:   "PHPSESSID",
			Value:  p.sessionId,
			Path:   "/",
			Domain: ".pixiv.net",
		},
	}
}

type TFanBoxer struct {
	tBaseAccount
}

func (p *TFanBoxer) Cookies() []*http.Cookie {
	return []*http.Cookie{
		{
			Name:   "FANBOXSESSID",
			Value:  p.sessionId,
			Path:   "/",
			Domain: ".fanbox.cc",
		},
	}
}

type IAccountV2 interface {
	Domain() string // fanbox、pixiver

	Uid() uint64

	State() uint32 // 账号状态

	Cookies() []*http.Cookie

	Json() map[string]any
}

type IAccountManager interface {
	Get(key string) (IAccountV2, error)
	Add(domain string, sessid string) (key string, err error)
	Del(key string) (IAccountV2, error)

	List() []map[string]any
}

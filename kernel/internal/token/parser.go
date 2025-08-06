package token

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YuzuWiki/Pixivlee/types"
)

const (
	PIXIV_SESSION_ID  = "PHPSESSID"
	FANBOX_SESSION_ID = "FANBOXSESSID"
)

const (
	_ int8 = iota
	BROWSE_MODE_VISITOR
	BROWSE_MODE_PIXIV
	BROWSE_MODE_FANBOX
)

func parser(mode int8, sid string) (*Token, error) {
	var (
		sessionId = strings.TrimSpace(sid)
		token     Token
	)
	switch mode {
	case BROWSE_MODE_VISITOR:
		return &token, nil

	case BROWSE_MODE_PIXIV, BROWSE_MODE_FANBOX:
		arr := strings.SplitN(sessionId, "_", 2)
		if len(arr) != 2 {
			return nil, fmt.Errorf("invalid sessionId, please check and try again.")
		}

		pid, err := strconv.ParseUint(arr[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid pid, please check and try again.")
		}

		if mode == BROWSE_MODE_PIXIV {
			return &Token{Pid: types.TPid(pid), SidPixiv: sessionId}, nil

		} else {
			return &Token{Pid: types.TPid(pid), SidFanbox: sessionId}, nil
		}
	default:
		return nil, fmt.Errorf(" Unknown Type")
	}
}

func New(cookies string) (*Token, error) {
	cookies = strings.TrimSpace(cookies)

	for _, field := range strings.Split(cookies, ";") {
		// eg: PHPSESSID=100_xxxxx
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}

		// pixiv sid
		if strings.HasPrefix(field, PIXIV_SESSION_ID) {
			return parser(BROWSE_MODE_PIXIV, field[len(PIXIV_SESSION_ID)+1:])
		}

		// fanbox sid
		if strings.HasPrefix(field, FANBOX_SESSION_ID) {
			return parser(BROWSE_MODE_FANBOX, field[len(FANBOX_SESSION_ID)+1:])
		}
	}
	return parser(BROWSE_MODE_VISITOR, "")
}

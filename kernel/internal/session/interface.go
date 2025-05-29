package session

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/YuzuWiki/Pixivlee/types"
)

type ISession interface {
	Mode() int8

	Pid() types.TPid
	SessId() string
}

type tSession struct {
	mode int8

	pid        types.TPid
	sid_pixiv  string
	sid_fanbox string
}

func (a *tSession) Mode() int8 {
	return a.mode
}

func (a *tSession) Pid() types.TPid {
	return a.pid
}

func (a *tSession) SessId() string {
	switch a.mode {
	case BROWSE_MODE_PIXIV:
		return a.sid_pixiv

	case BROWSE_MODE_FANBOX:
		return a.sid_fanbox

	default:
		return ""
	}
}

func New(mode int8, sessionId string) (ISession, error) {
	sessionId = strings.TrimSpace(sessionId)

	switch mode {
	case BROWSE_MODE_VISITOR:
		return &tSession{mode: BROWSE_MODE_VISITOR}, nil

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
			return &tSession{mode: mode, pid: types.TPid(pid), sid_pixiv: sessionId}, nil

		} else {
			return &tSession{mode: mode, pid: types.TPid(pid), sid_fanbox: sessionId}, nil
		}

	default:
		return nil, fmt.Errorf(" Unknown Session Type")
	}
}

package auth

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/YuzuWiki/Pixivlee/types"
)

func newToken(mode int8, sessionId string) (ISessId, error) {
	sessionId = strings.TrimSpace(sessionId)

	arr := strings.SplitN(strings.TrimSpace(sessionId), "_", 2)
	if len(arr) != 2 {
		return nil, errors.Wrap(ErrAuth, "parsing sessid failed, please check and try again.")
	}

	pid, err := strconv.ParseUint(arr[0], 10, 64)
	if err != nil {
		return nil, errors.Wrap(ErrAuth, "invalid session pid, please check and try again.")
	}
	return &tSessId{mode: mode, pid: types.TPid(pid), sessionId: sessionId}, nil
}

func NewVisitor() (ISessId, error) {
	return &tSessId{mode: BROWSE_MODE_VISITOR}, nil
}

func NewPixiv(sessionId string) (ISessId, error) {
	return newToken(BROWSE_MODE_PIXIV, sessionId)
}

func NewFanbox(sessionId string) (ISessId, error) {
	return newToken(BROWSE_MODE_FANBOX, sessionId)
}

// ParsecCookie default use visitor
func ParsecCookie(cookies string) (ISessId, error) {
	cookies = strings.TrimSpace(cookies)
	for _, field := range strings.Split(cookies, ";") {
		field = strings.TrimSpace(field)

		if strings.HasPrefix(field, PIXIV_SESSION_ID) {
			return NewPixiv(field[len(PIXIV_SESSION_ID)+1:])
		}

		if strings.HasPrefix(field, FANBOX_SESSION_ID) {
			return NewFanbox(field[len(FANBOX_SESSION_ID)+1:])
		}
	}
	return NewVisitor()
}

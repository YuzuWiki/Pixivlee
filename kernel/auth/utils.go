package auth

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"

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

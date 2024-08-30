package types

import (
	"fmt"
	"strconv"
	"strings"

	resty "github.com/go-resty/resty/v2"
)

type IKernel interface {
	NewRequests() *resty.Request
}

type IPixiver interface {
	Pid() TPid
	SessionID() string
}

type TPixiver struct {
	pid       TPid
	sessionId string
}

func (p *TPixiver) Pid() TPid {
	return p.pid
}

func (p *TPixiver) SessionID() string {
	return p.sessionId
}

func NewPixiver(sessionId string) (IPixiver, error) {
	sessionId = strings.TrimSpace(sessionId)
	if len(sessionId) == 0 {
		return nil, fmt.Errorf("sessionid is invalid, must not be empty")
	}

	pixiver := &TPixiver{
		sessionId: sessionId,
	}

	arr := strings.SplitN(strings.TrimSpace(sessionId), "_", 2)
	if len(arr) != 2 {
		return nil, fmt.Errorf("sessionid is invalid, format error")
	}

	pid, err := strconv.ParseUint(arr[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("sessionid is invalid, (pid) %s", err.Error())
	}
	pixiver.pid = TPid(pid)

	return pixiver, nil
}

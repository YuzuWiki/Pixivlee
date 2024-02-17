package Pixivlee

import (
	"strconv"
	"strings"
	"time"
)

type Pixiver struct {
	pid       TPid
	sessionID string

	state   int
	limitAt int
}

func (p *Pixiver) Pid() TPid {
	if p.pid > 0 {
		return p.pid
	}

	if arr := strings.SplitN(p.sessionID, "_", 2); len(arr) == 2 {
		pid, _ := strconv.ParseUint(arr[0], 10, 64)
		return TPid(pid)
	}

	return 0
}

func (p *Pixiver) SessionID() string {
	return p.sessionID
}

func (p *Pixiver) State() int {
	return p.state
}

func (p *Pixiver) SetPid(pid TPid) {
	if p.Pid() == 0 {
		p.pid = pid
	}
}

func (p *Pixiver) UpdateState(state int) {
	if p.State() == state {
		return
	}

	switch state {
	case PixiverNormal:
		p.state = PixiverNormal
		p.limitAt = 0

	case PixiverRateLimiting:
		p.state = PixiverRateLimiting
		p.limitAt = time.Now().Nanosecond()

	case PixiverInvalid:
		p.state = PixiverInvalid
		p.limitAt = 0
	}
}

func NewPixiver(sessionID string) IPixiver {
	return &Pixiver{sessionID: sessionID}
}

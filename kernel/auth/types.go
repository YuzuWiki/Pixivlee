package auth

import (
	"github.com/YuzuWiki/Pixivlee/types"
)

type ISessId interface {
	Mode() int8

	Pid() types.TPid
	SessId() string
}

type tSessId struct {
	mode int8

	pid       types.TPid
	sessionId string
}

func (a *tSessId) Mode() int8 {
	return a.mode
}

func (a *tSessId) Pid() types.TPid {
	return a.pid
}

func (a *tSessId) SessId() string {
	return a.sessionId
}

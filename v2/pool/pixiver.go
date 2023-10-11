package pool

import v2 "github.com/YuzuWiki/Pixivlee/v2"

type Pixiver struct {
	pid       v2.TPid
	status    uint8
	sessionId string
}

func (p *Pixiver) Pid() v2.TPid {
	return 0
}

func (p *Pixiver) SessID() string {
	return p.sessionId
}

func (p *Pixiver) IsEnable() bool {
	return false
}

type NewPixiver struct {
}

func (p NewPixiver) SessID(sessId string) (error, v2.IPixiver) {
	return nil, nil
}

func (p NewPixiver) Auth(username string, password string) (error, v2.IPixiver) {
	return nil, nil
}

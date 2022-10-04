package Pixivlee

import "github.com/YuzuWiki/Pixivlee/common"

type Context struct {
	phpSessID string
}

func (ctx *Context) PhpSessID() string {
	return ctx.phpSessID
}

func NewContext(phpSessId string) common.IContext {
	return &Context{phpSessID: phpSessId}
}

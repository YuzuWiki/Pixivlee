package Pixivlee

import "github.com/YuzuWiki/Pixivlee/common"

var pool common.IPool

func init() {
	pool = newPool()
}

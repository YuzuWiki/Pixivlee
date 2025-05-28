package apis

import (
	"github.com/YuzuWiki/Pixivlee/apis/internal/fanbox"
	"github.com/YuzuWiki/Pixivlee/apis/internal/pixiv"
)

var (
	// Register fanbox api
	RegisterFanbox = fanbox.RegisterProvider

	// Register fanbox api
	RegisterPixiv = pixiv.RegisterProvider
)

package Pixivlee

import (
	"fmt"
	"testing"

	"github.com/YuzuWiki/Pixivlee/kernel"
)

func Test_NewPixiv(t *testing.T) {
	options := kernel.Options{
		Http: kernel.HttpOption{
			IsDebug:  false,
			Timeout:  10,
			ProxyUrl: "socks5://127.0.0.1:27005",
		},
		Pixiv: kernel.PixivOption{
			PhpSessId: "97638218_dXVADfTGL2Fn31QgYn8nRnGxqpp80sci",
		},
	}

	app, err := NewFanbox(options)
	if err != nil {
		t.Error(err)
		return
	}

	data, err := app.Post.Info(6063594)
	//data, err := app.Creator.Creator("zhibujiloom")
	//data, err := app.FanBox.Creator("zhibujiloom")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", data)
}

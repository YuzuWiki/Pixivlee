package Pixivlee

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

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

	Iterator := app.Post.List("zhibujiloom", 10)
	for page, data := range Iterator.Next() {

		d, _ := json.Marshal(data)

		fmt.Println(string(d))
		fmt.Printf("===================== %d =====================", page)
		time.Sleep(2 * time.Second)
	}

	//data, err := app.Creator.Get(13695413)
	//data, err := app.Creator.Creator("zhibujiloom")
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//fmt.Printf("%+v\n", data)
}

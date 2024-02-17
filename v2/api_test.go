package v2

import (
	"fmt"
	"testing"
)

func TestPixivApi_AccountInfo(t *testing.T) {
	SetProxy("http://127.0.0.1:27002")
	p := NewPixiver("85056881_5RztYakS3S7I01wGTb3JVHdbYQ8NdDuY")

	api := PixivApi{}
	//data, err := api.BookmarkShow(p, 50783209, "", 9, 10)
	data, err := api.RankIllust(p, RankDailyR18, 0, "")
	if err != nil {
		fmt.Println("error: ", err.Error())
	} else {
		fmt.Println(fmt.Sprintf("success: %+v", data))
	}
}

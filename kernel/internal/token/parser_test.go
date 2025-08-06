package token

import (
	"os"
	"testing"

	_ "github.com/YuzuWiki/Pixivlee/testdata"
)

func Test_New(t *testing.T) {
	sessId, cookies := os.Getenv("PIXIV_SESSID"), os.Getenv("PIXIV_COOKIE")

	a, err := New(cookies)
	if err != nil {
		t.Error(err)
		return
	}

	if a.SidPixiv != sessId {
		t.Errorf("parse error (not equal)")
		return
	}
	return
}

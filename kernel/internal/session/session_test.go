package session

import (
	"os"
	"testing"

	_ "github.com/YuzuWiki/Pixivlee/testdata"
)

func Test_TAuthCookie(t *testing.T) {
	sessId, cookies := os.Getenv("PIXIV_SESSID"), os.Getenv("PIXIV_COOKIE")

	a, err := Session(cookies)
	if err != nil {
		t.Error(err)
		return
	}

	if a.SessId() != sessId {
		t.Errorf("parse error (not equal)")
		return
	}
	return
}

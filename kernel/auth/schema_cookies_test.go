package auth

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_TAuthCookie(t *testing.T) {
	err := godotenv.Load("../../testdata/config.env")
	if err != nil {
		t.Error(err)
		return
	}
	sessId, cookies := os.Getenv("PIXIV_SESSID"), os.Getenv("PIXIV_COOKIE")

	a, err := NewAuthCookie(cookies)
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

package requests

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func (r *requests) ensureJar() {
	if r.Jar == nil {
		r.Jar, _ = cookiejar.New(nil)
	}
}

func (r *requests) SetCookies(rawURL string, cookies ...*http.Cookie) error {
	if len(rawURL) == 0 || len(cookies) == 0 {
		return fmt.Errorf("invalid params")
	}

	r.ensureJar()
	if !strings.HasPrefix(rawURL, "http") {
		rawURL = "https://" + rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	r.Jar.SetCookies(u, cookies)
	return nil
}

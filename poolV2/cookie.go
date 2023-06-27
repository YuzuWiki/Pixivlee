package poolV2

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type CookieProxy struct {
	jar *cookiejar.Jar
}

func (p *CookieProxy) Set(rawUrl string, cookies ...*http.Cookie) error {
	if len(rawUrl) == 0 || len(cookies) == 0 {
		return fmt.Errorf("invalid cookie")
	}

	u, err := url.Parse(rawUrl)
	if err != nil {
		return err
	}

	p.jar.SetCookies(u, cookies)
	return nil
}

func (p *CookieProxy) Get(rawUrl string, name string) (*http.Cookie, error) {
	if len(rawUrl) == 0 || len(name) == 0 {
		return nil, fmt.Errorf("invalid params")
	}

	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	for _, c := range p.jar.Cookies(u) {
		cookie := c
		if c.Name == name {
			return cookie, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

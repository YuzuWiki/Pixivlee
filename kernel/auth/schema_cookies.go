package auth

import (
	"strings"
)

// ParsecCookie default use visitor
func NewAuthCookie(cookies string) (ISessId, error) {
	cookies = strings.TrimSpace(cookies)
	for _, field := range strings.Split(cookies, ";") {
		field = strings.TrimSpace(field)

		if strings.HasPrefix(field, PIXIV_SESSION_ID) {
			return NewPixiv(field[len(PIXIV_SESSION_ID)+1:])
		}

		if strings.HasPrefix(field, FANBOX_SESSION_ID) {
			return NewFanbox(field[len(FANBOX_SESSION_ID)+1:])
		}
	}
	return NewVisitor()
}

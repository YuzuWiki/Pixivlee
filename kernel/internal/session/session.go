package session

import (
	"strings"
)

// Session parse cookie, default visitor
func Session(cookies string) (ISession, error) {
	cookies = strings.TrimSpace(cookies)
	for _, field := range strings.Split(cookies, ";") {
		field = strings.TrimSpace(field)

		if strings.HasPrefix(field, PIXIV_SESSION_ID) {
			return New(BROWSE_MODE_PIXIV, field[len(PIXIV_SESSION_ID)+1:])
		}

		if strings.HasPrefix(field, FANBOX_SESSION_ID) {
			return New(BROWSE_MODE_FANBOX, field[len(FANBOX_SESSION_ID)+1:])
		}
	}
	return New(BROWSE_MODE_VISITOR, "")
}

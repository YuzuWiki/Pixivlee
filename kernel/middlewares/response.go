package middlewares

import (
	"fmt"
	"github.com/YuzuWiki/Pixivlee/kernel/httpx"
)

func DefaultResponse() func(iClient httpx.ISession, iResponse httpx.IResponse) error {
	return func(iClient httpx.ISession, iResponse httpx.IResponse) error {
		if HttpCode := iResponse.StatusCode(); HttpCode != 200 {
			body := iResponse.Body()

			// eg. 400  {"error":"general_error"}
			return fmt.Errorf(fmt.Sprintf("%d  %s", HttpCode, string(body)))
		}
		return nil
	}
}

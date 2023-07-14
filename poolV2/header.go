package poolV2

import (
	"net/http"

	"github.com/YuzuWiki/Pixivlee"
)

type Header struct {
	header http.Header
}

func (h *Header) Set(headers ...Pixivlee.THeader) error {
	for idx := range headers {
		h.header.Set(headers[idx].Key, headers[idx].Value)
	}
	return nil
}

func (h *Header) Get(key string) string {
	return h.header.Get(key)
}

package options

import (
	"fmt"
	"strings"
)

func formatUrl(host string) string {
	host = strings.TrimSpace(host)
	if host == "" {
		return ""
	}

	if strings.HasPrefix(host, "http") {
		return host
	}
	return fmt.Sprintf("https://%s", host)
}

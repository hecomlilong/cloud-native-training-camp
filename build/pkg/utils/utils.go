package utils

import (
	"net/http"
	"strings"
)

func FindClientIp(r *http.Request) string {
	if r == nil {
		return ""
	}
	items := strings.Split(r.RemoteAddr, ":")
	if len(items) > 0 {
		return items[0]
	}
	return ""
}

package utils

import (
	"math/rand"
	"net/http"
	"strings"
	"time"
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

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

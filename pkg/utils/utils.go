package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
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
func ForwardHandler(w http.ResponseWriter, r *http.Request, serviceName string) error {
	tag := fmt.Sprintf("%s:ForwardHandler", serviceName)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", serviceName), nil)
	if err != nil {
		GlogWrapper.Warningf("%s:err:%v", tag, err)
		return errors.Wrap(err, tag)
	}
	lowerCaseHeader := make(http.Header)
	for k, items := range r.Header {
		lowerCaseHeader[strings.ToLower(k)] = items
	}
	GlogWrapper.Infof("%s:header:%v", tag, lowerCaseHeader)
	req.Header = lowerCaseHeader
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		GlogWrapper.Warningf("%s:client.Do:err:%v", tag, err)
		return errors.Wrap(err, tag)
	}
	if resp != nil {
		resp.Write(w)
	}
	GlogWrapper.Infof("%s success", tag)
	return nil
}

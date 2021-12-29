package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/hecomlilong/cloud-native-training-camp/pkg/utils"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	s := &http.Server{
		Addr:           ":1880",
		Handler:        new(Handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	glog.V(2).Info("Starting http server...")
	log.Fatal(s.ListenAndServe())
}

type Handler struct{}

func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if w == nil {
		glog.Warning("Handler.ServeHTTP got nil ResponseWriter")
		return
	}
	if r == nil {
		glog.Warning("Handler.ServeHTTP got nil Request")
	}
	p.HeaderHandler(w, r)
	p.VersionHandler(w, r)
	status := p.HealthzHandler(w, r)
	p.StdoutHandler(w, r, status)
}

// 接收客户端 request，并将 request 中带的 header 写入 response header
func (p *Handler) HeaderHandler(w http.ResponseWriter, r *http.Request) {
	for k, items := range r.Header {
		for _, item := range items {
			w.Header().Add(k, item)
		}
	}
}

// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
func (p *Handler) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Version", os.Getenv("VERSION"))
}

// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
func (p *Handler) StdoutHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	glog.V(2).Infof("client ip address:%s", utils.FindClientIp(r))
	glog.V(2).Infof("http code:%d", statusCode)
}

// 当访问 localhost/healthz 时，应返回 200
func (p *Handler) HealthzHandler(w http.ResponseWriter, r *http.Request) (statusCode int) {
	if r.URL.Path == "/healthz" {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusAccepted
	}
	w.WriteHeader(statusCode)
	return
}

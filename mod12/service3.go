package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/hecomlilong/cloud-native-training-camp/pkg/metrics"
	"github.com/hecomlilong/cloud-native-training-camp/pkg/utils"
	configObj "github.com/hecomlilong/cloud-native-training-camp/pkg/utils/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var config = flag.String("config", "", "config as uri format")

func main() {
	flag.Parse()
	err := configObj.InitConfig(*config)
	if err != nil {
		fmt.Printf("config err:%v\n", err)
		return
	}
	defer glog.Flush()
	s := &http.Server{
		Addr:           ":1880",
		Handler:        new(Handler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utils.GlogWrapper.Info("Starting http server...")
	metrics.Register()

	var graceExit string
	errChan := make(chan error)
	exitChan := make(chan string)
	go func() {
		err = s.ListenAndServe()
		errChan <- err
	}()
	// 创建一个os.Signal channel
	sigs := make(chan os.Signal, 1)
	//注册要接收的信号，syscall.SIGINT:接收ctrl+c ,syscall.SIGTERM:程序退出
	//信号没有信号参数表示接收所有的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//此goroutine为执行阻塞接收信号。一旦有了它，它就会打印出来。
	//然后通知程序可以完成。
	go func() {
		sig := <-sigs
		utils.GlogWrapper.Warningf("get signal:%v", sig)
		exitChan <- "gracefully exit"
	}()

	select {
	case err = <-errChan:
		utils.GlogWrapper.Warningf("server err:%v", err)
		return
	case graceExit = <-exitChan:
		utils.GlogWrapper.Infof("msg:%v", graceExit)
		return
	}
}

type Handler struct{}

func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if w == nil {
		utils.GlogWrapper.Warning("Handler.ServeHTTP got nil ResponseWriter")
		return
	}
	if r == nil {
		utils.GlogWrapper.Warning("Handler.ServeHTTP got nil Request")
	}
	p.HeaderHandler(w, r)
	p.VersionHandler(w, r)
	status := p.HealthzHandler(w, r)
	p.StdoutHandler(w, r, status)
	p.RootHandler(w, r)
	p.MatricsHandler(w, r)
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
	utils.GlogWrapper.Infof("client ip address:%s", utils.FindClientIp(r))
	utils.GlogWrapper.Infof("http code:%d", statusCode)
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

func (p *Handler) MatricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/metrics" {
		promhttp.Handler()
	}
}

func (p *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		return
	}
	utils.GlogWrapper.Info("entering root handler")
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	visitor := r.URL.Query().Get("visitor")
	delay := utils.RandInt(0, 2000)
	time.Sleep(time.Millisecond * time.Duration(delay))
	if visitor != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", visitor))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	utils.GlogWrapper.Infof("Respond in %d ms", delay)
}

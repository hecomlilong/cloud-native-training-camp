package utils

import (
	"github.com/golang/glog"
	configObj "github.com/hecomlilong/cloud-native-training-camp/pkg/utils/config"
)

type Glog struct{}

var GlogWrapper Glog

func LogLevel(l glog.Level) bool {
	return int64(l) >= configObj.GetConfig().Int64("glog_level")
}

func (v Glog) Debug(args ...interface{}) {
	if bool(glog.V(DEBUG)) || LogLevel(DEBUG) {
		glog.Info(args...)
	}
}

func (v Glog) Debugln(args ...interface{}) {
	if bool(glog.V(DEBUG)) || LogLevel(DEBUG) {
		glog.Infoln(args...)
	}
}

func (v Glog) Debugf(format string, args ...interface{}) {
	if bool(glog.V(DEBUG)) || LogLevel(DEBUG) {
		glog.Infof(format, args...)
	}
}

func (v Glog) Info(args ...interface{}) {
	if bool(glog.V(INFO)) || LogLevel(INFO) {
		glog.Info(args...)
	}
}

func (v Glog) Infoln(args ...interface{}) {
	if bool(glog.V(INFO)) || LogLevel(INFO) {
		glog.Infoln(args...)
	}
}

func (v Glog) Infof(format string, args ...interface{}) {
	if bool(glog.V(INFO)) || LogLevel(INFO) {
		glog.Infof(format, args...)
	}
}

func (v Glog) Warning(args ...interface{}) {
	if bool(glog.V(WARN)) || LogLevel(WARN) {
		glog.Info(args...)
	}
}

func (v Glog) Warningln(args ...interface{}) {
	if bool(glog.V(WARN)) || LogLevel(WARN) {
		glog.Infoln(args...)
	}
}

func (v Glog) Warningf(format string, args ...interface{}) {
	if bool(glog.V(WARN)) || LogLevel(WARN) {
		glog.Infof(format, args...)
	}
}

func (v Glog) Error(args ...interface{}) {
	if bool(glog.V(ERROR)) || LogLevel(ERROR) {
		glog.Info(args...)
	}
}

func (v Glog) Errorln(args ...interface{}) {
	if bool(glog.V(ERROR)) || LogLevel(ERROR) {
		glog.Infoln(args...)
	}
}

func (v Glog) Errorf(format string, args ...interface{}) {
	if bool(glog.V(ERROR)) || LogLevel(ERROR) {
		glog.Infof(format, args...)
	}
}

func (v Glog) Fatal(args ...interface{}) {
	if bool(glog.V(FATAL)) || LogLevel(FATAL) {
		glog.Info(args...)
	}
}

func (v Glog) Fatalln(args ...interface{}) {
	if bool(glog.V(FATAL)) || LogLevel(FATAL) {
		glog.Infoln(args...)
	}
}

func (v Glog) Fatalf(format string, args ...interface{}) {
	if bool(glog.V(FATAL)) || LogLevel(FATAL) {
		glog.Infof(format, args...)
	}
}

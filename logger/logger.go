package logger

import (
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(1000)
	log.SetLogger("console", "")
	log.SetLogger("file", `{"filename":"./log/copy.log"}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
}
func Debug(f string, v ...interface{}) {
	log.Debug(f, v...)
}

func Error(f string, v ...interface{}) {
	log.Error(f, v...)
}
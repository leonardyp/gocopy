package logger

import (
	"log"
	"os"
	"runtime"
)

type ConsoleLogger struct {
}

var cl = ConsoleLogger{}

func (this *ConsoleLogger) Debug(format string, v ...interface{}) {
	level = DEBUG
	std.Println(colors[DEBUG]("console", format, v...))
}
func (this *ConsoleLogger) Error(format string, v ...interface{}) {
	level = ERROR
	std.Println(colors[ERROR]("console", format, v...))
}
func (this *ConsoleLogger) DebugFunCall(format string, v ...interface{}) {
	for i := 0; i < DEBUG_DEPTH; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			std.Printf(blue("[%v:%v")+magenta(",func:%v]"), file, line, runtime.FuncForPC(funcName).Name())
		}
	}
}
func StdDebug(format string, v ...interface{}) {
	cl.Debug(format, v...)
}
func StdError(format string, v ...interface{}) {
	cl.Error(format, v...)
}
func StdDebugFunCall(format string, v ...interface{}) {
	cl.DebugFunCall(format, v...)
}
func (this *ConsoleLogger) SetLoggerConsole() {
	if _, ok := destOut["console"]; !ok {
		destOut["console"] = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	}
}

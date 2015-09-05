package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

const (
	EMERGENCY = iota
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFORMATION
	DEBUG
	DEBUG_DEPTH = 8
	ERROR_DEPTH
)

var colors = []Brush{
	NewBrush("1;37"), // Emergency	white
	NewBrush("1;36"), // Alert			cyan
	NewBrush("1;35"), // Critical   magenta
	NewBrush("1;31"), // Error      red
	NewBrush("1;33"), // Warning    yellow
	NewBrush("1;32"), // Notice			green
	NewBrush("1;34"), // Informational	blue
	NewBrush("1;34"), // Debug      blue
}

type Brush func(string) string

func NewBrush(color string) Brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}
func Debug(format string, v ...interface{}) {
	l.Println(colors[DEBUG](fmt.Sprintf(format, v...)))
}
func Error(format string, v ...interface{}) {
	l.Println(colors[ERROR](fmt.Sprintf(format, v...)))
}
func DebugDetail(format string, v ...interface{}) {
	Debug(format, v...)
	for i := 0; i < DEBUG_DEPTH; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			Debug("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		}
	}
}

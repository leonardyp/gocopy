package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

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
var destOut = map[string]*log.Logger{}

var (
	std *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

var level int

func magenta(content string) string {
	return "\033[1;35m" + content + "\033[0m"
}
func white(content string) string {
	return "\033[1;37m" + content + "\033[0m"
}
func red(content string) string {
	return "\033[1;31m" + content + "\033[0m"
}
func blue(content string) string {
	return "\033[1;34m" + content + "\033[0m"
}

type Brush func(dest, format string, v ...interface{}) string

func NewBrush(color string) Brush {
	pre := "\033["
	reset := "\033[0m"
	return func(dest, format string, v ...interface{}) string {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			var gopaths []string

			if runtime.GOOS == `darwin` {
				gopaths = strings.Split(os.Getenv("GOPATH"), ":")
			}
			if runtime.GOOS == `windows` {
				gopaths = strings.Split(os.Getenv("GOPATH"), ";")
			}
			for _, gopath := range gopaths {
				if strings.Contains(file, gopath) {
					file = strings.TrimPrefix(file, gopath+"/src/")
				}
			}
			if level == DEBUG {
				format = "[" + file + ":" + strconv.Itoa(line) + "] [D] " + format
			} else if level == ERROR {
				format = "[" + file + ":" + strconv.Itoa(line) + "] [E] " + format
			}
		}
		if dest == "console" {
			return pre + color + "m" + fmt.Sprintf(format, v...) + reset
		}
		return fmt.Sprintf(format, v...)
	}
}

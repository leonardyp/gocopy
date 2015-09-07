package logger

import (
	"log"
	"os"
	"runtime"
	"strings"
)

type FileLogger struct {
	FileName string
	fd       *os.File
}

var fl = FileLogger{}

func (this *FileLogger) Debug(format string, v ...interface{}) {
	level = DEBUG
	destOut["file"].Println(colors[DEBUG]("file", format, v...))
}
func (this *FileLogger) Error(format string, v ...interface{}) {
	level = ERROR

	destOut["file"].Println(colors[ERROR]("file", format, v...))
}
func (this *FileLogger) DebugFunCall(format string, v ...interface{}) {
	for i := 0; i < DEBUG_DEPTH; i++ {
		funcName, file, line, ok := runtime.Caller(i)
		if ok {
			destOut["file"].Printf("[%v:%v] [D] func:%v]", file, line, runtime.FuncForPC(funcName).Name())
		}
	}
}
func (this *FileLogger) Init() error {
	fd, err := this.createLogFile()
	if err != nil {
		return err
	}
	this.fd = fd
	this.SetLoggerFile(fd)
	return err
}
func (this *FileLogger) createLogFile() (*os.File, error) {
	fd, err := os.OpenFile(this.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	return fd, err
}
func (this *FileLogger) SetLoggerFile(fd *os.File) {
	destOut["file"] = log.New(fd, "", log.Ldate|log.Ltime)
}
func StartFileLogger(fileName string) {
	if strings.TrimSpace(fileName) == "" {
		fl.FileName = "log.log"
	} else {
		fl.FileName = fileName
	}
	fl.Init()
}

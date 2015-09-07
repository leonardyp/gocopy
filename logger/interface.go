package logger

type LoggerInterface interface {
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	DebugFunCall(format string, v ...interface{})
}

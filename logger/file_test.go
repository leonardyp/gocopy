package logger

import "testing"

func TestFileDebug(t *testing.T) {
	StartFileLogger("")
	fl.Debug("%#v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

func TestFileError(t *testing.T) {
	StartFileLogger("demo.log")
	fl.Error("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

func TestFileDebugDetail(t *testing.T) {
	StartFileLogger("log.log")
	fl.DebugFunCall("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

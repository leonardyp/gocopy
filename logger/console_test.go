package logger

import (
	"testing"
)

type Demo struct {
	Name  string
	Money float32
}

func TestDebug(t *testing.T) {
	var cl = ConsoleLogger{}
	cl.Debug("%#v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

func TestError(t *testing.T) {
	var cl = ConsoleLogger{}
	cl.Error("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

func TestDebugDetail(t *testing.T) {
	var cl = ConsoleLogger{}
	cl.DebugFunCall("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

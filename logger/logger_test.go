package logger

import (
	"testing"
)

type Demo struct {
	Name  string
	Money float32
}

func TestDebug(t *testing.T) {
	Debug("%#v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

func TestError(t *testing.T) {
	Error("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

func TestDebugDetail(t *testing.T) {
	DebugFunCall("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

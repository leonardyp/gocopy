package logger

import (
	"testing"
)

type Demo struct {
	Name  string
	Money float32
}

func TestDebug(t *testing.T) {
	StdDebug("%#v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

func TestError(t *testing.T) {
	StdError("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

func TestDebugDetail(t *testing.T) {
	StdDebugFunCall("%#v:%s", Demo{
		"leonard",
		23.4,
	}, "hello world")
}

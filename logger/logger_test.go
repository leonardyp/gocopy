package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("%+v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

func TestDebugDetail(t *testing.T) {
	DebugDetail("%+v:%s", struct {
		Name  string
		Money float32
	}{
		"leonard",
		23.4,
	}, "hello world")
}

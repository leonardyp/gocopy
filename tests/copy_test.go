package tests

import (
	"github.com/leonardyp/gocopy"
	"github.com/leonardyp/gocopy/logger"
	"testing"
)

type A struct {
	Id    int64
	Name  *string
	Sex   bool
	Uid   int64
	Info  string
	Price float32
}
type B struct {
	Mid  int `cp:"Id"`
	Name string
	Sex  bool
}

func TestCopy(t *testing.T) {
	var str = "demoA"
	a := &A{
		Name:  &str,
		Id:    1,
		Uid:   320,
		Price: 12.3,
		Info:  "info...",
		Sex:   true,
	}
	b := &B{Name: "demo"}
	logger.StdDebug("before copy :%v:%v:%v", b.Mid, b.Name, b.Sex)
	gocopy.Copy(a, b)
	logger.StdDebug("after copy :%v:%v:%v", b.Mid, b.Name, b.Sex)
}

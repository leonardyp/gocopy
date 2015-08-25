package tests

import (
	"github.com/leonardyp/gocopy"
	"testing"
)

type A struct {
	Id    int64
	Name  *string
	sex   bool
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
	var str = "demo"
	a := &A{
		Name:  &str,
		Id:    1,
		Uid:   320,
		Price: 12.3,
		Info:  "info...",
	}
	b := &B{Name: "demo"}
	gocopy.Copy(a, b)
}

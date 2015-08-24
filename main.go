package main

import "copy"

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

func main() {

	var str = "demo"
	a := &A{
		Name:  &str,
		Id:    1,
		Uid:   320,
		Price: 12.3,
		Info:  "info...",
	}
	b := &B{Mid: 2}
	copy.Copy(a, b)
}

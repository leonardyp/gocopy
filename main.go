package main

import (
	"github.com/lunny/log"
	"reflect"
)

type A struct {
	Id    int64
	Name  *string
	sex   bool
	Uid   uint64
	Info  string
	Price float32
}
type B struct {
	Mid  int `cp:"Id"`
	Name string
	sex  bool
}

var ma, mb = map[string]reflect.Value{}, map[string]reflect.Value{}

func main() {

	var str = "demo"
	a := A{
		Name: &str,
		Id:   1,
		Uid:  320,
	}
	rta := reflect.TypeOf(a)
	rva := reflect.ValueOf(a)
	for i := 0; i < rta.NumField(); i++ {
		ma[rta.Field(i).Name] = rva.Field(i)
	}

	for k, v := range ma {
		switch v.Kind() {
		case reflect.String:
			log.Error("%s:%#v", k, v.String())
		case reflect.PtrTo(reflect.TypeOf("")).Kind():
			log.Error("%s:%#v", k, v.Elem().String())
		case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
			log.Error("%s:%#v", k, v.Int())
		case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
			log.Error("%s:%#v", k, v.Uint())
		case reflect.Float32, reflect.Float64:
			log.Error("%s:%#v", k, v.Float())
		case reflect.Bool:
			log.Error("%s:%#v", k, v.Bool())
		default:
			log.Error("%s:%v", k, v.Type())
		}
	}

	b := &B{Mid: 2}
	rtb := reflect.TypeOf(b)
	rvb := reflect.ValueOf(b)
	for i := 0; i < rtb.NumField(); i++ {
		if rtb.Field(i).Tag.Get("cp") == "" {
			mb[rtb.Field(i).Name] = rvb.Field(i)
		} else {
			mb[rtb.Field(i).Tag.Get("cp")] = rvb.Field(i)
		}
	}

}

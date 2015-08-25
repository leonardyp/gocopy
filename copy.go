package gocopy

import (
	"github.com/leonardyp/gocopy/logger"
	"reflect"
)

var ma, mb = map[string]reflect.Value{}, map[string]reflect.Value{}

func Copy(a, b interface{}) {
	//todo defer
	rta := reflect.TypeOf(a).Elem()
	if rta.Kind() != reflect.Struct {
		panic("request struct accept " + rta.Kind().String())
	}

	rtb := reflect.TypeOf(b).Elem()
	if rtb.Kind() != reflect.Struct {
		panic("request struct accept " + rta.Kind().String())
	}

	rva := reflect.ValueOf(a).Elem()
	rvb := reflect.ValueOf(b).Elem()

	for i := 0; i < rta.NumField(); i++ {
		v := rva.Field(i)
		switch rva.Field(i).Kind() {
		case reflect.String:
			ma[rta.Field(i).Name] = rva.Field(i)

		case reflect.PtrTo(reflect.TypeOf("")).Kind():
			ma[rta.Field(i).Name] = rva.Field(i)

		case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
			ma[rta.Field(i).Name] = rva.Field(i)

		case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
			ma[rta.Field(i).Name] = rva.Field(i)

		case reflect.Float32, reflect.Float64:
			ma[rta.Field(i).Name] = rva.Field(i)

		case reflect.Bool:
			ma[rta.Field(i).Name] = rva.Field(i)

		default:
			logger.Debug("%v", v.Type())
		}
	}
	var nb = map[string]string{}
	for i := 0; i < rtb.NumField(); i++ {
		if rtb.Field(i).Tag.Get("cp") == "" {
			mb[rtb.Field(i).Name] = rvb.Field(i)
		} else {
			nb[rtb.Field(i).Name] = rtb.Field(i).Tag.Get("cp")
			mb[rtb.Field(i).Name] = rvb.Field(i)
		}
	}

	for k, _ := range mb {
		old := k
		if nameB, ok := nb[k]; ok {
			k = nameB
		}
		if _, ok := ma[k]; ok {
			if _, ok = mb[old]; ok {
				if ma[k].Kind() == reflect.Ptr {

					if ma[k].Elem().Kind() == reflect.String {
						mb[old].SetString(ma[k].String())
					}
					if ma[k].Elem().Kind() == reflect.Bool {
						mb[old].SetBool(true)
					}
					if ma[k].Elem().Kind() == reflect.Int {
						mb[old].SetInt(ma[k].Int())
					}

				} else {

					if ma[k].Kind() == reflect.String {
						mb[old].SetString(ma[k].String())
					}
					if ma[k].Kind() == reflect.Bool {
						mb[old].SetBool(true)
					}
					if ma[k].Kind() == reflect.Int64 {
						mb[old].SetInt(ma[k].Int())
					}
				}
			}
		}
	}
	logger.Debug("%#v", b)
}

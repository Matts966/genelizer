package reflectcan

import (
	"reflect"
)

func test1(i *interface{}) {
	rv := reflect.ValueOf(i)
	rv.Addr() // want `should call CanAddr before calling Addr`
}

func test2(i *interface{}) {
	rv := reflect.ValueOf(i)
	if !rv.CanAddr() {
		return
	}
	rv.Addr()
}

func test3(i *interface{}) {
	rv := reflect.ValueOf(i)
	if rv.CanAddr() {
		rv.Addr()
	}
}

func test4(i *interface{}) {
	rv := reflect.ValueOf(i)
	rv.Interface() // want `should call CanInterface before calling Interface`
}

func test5(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 reflect.Value
	rv.Set(rv2) // want `should call CanSet before calling Set`
}

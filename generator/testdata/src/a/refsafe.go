package a

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

func test1(i *interface{}) {
	rv := reflect.ValueOf(i)
	rv.Addr() // want `CanAddr should be called before calling Addr`
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
	rv.Interface() // want `CanInterface should be called before calling Interface`
}

func test5(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 reflect.Value
	rv.Set(rv2) // want `CanSet should be called before calling Set`
}

func test6(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 unsafe.Pointer
	if rv.CanSet() {
		rv.SetPointer(rv2) // want `Kind should be UnsafePointer when calling SetPointer`
	}
}

func test7(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 unsafe.Pointer
	if rv.CanSet() && rv.Kind() == reflect.Interface {
		rv.SetPointer(rv2) // want `Kind should be UnsafePointer when calling SetPointer`
	}
}

func test8(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 unsafe.Pointer
	if rv.CanSet() && rv.Kind() == reflect.UnsafePointer {
		rv.SetPointer(rv2)
	}
}

func test9(i *interface{}) {
	rv := reflect.ValueOf(i)
	var rv2 string
	if rv.CanSet() && rv.Kind() == reflect.UnsafePointer {
		rv.SetString(rv2) // want `Kind should be String when calling SetString`
	}
}

func set(field reflect.Value, s string) error {
	switch field.Kind() {
	default:
		return fmt.Errorf("unknown type")
	case reflect.String:
		field.SetString(s)
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		field.SetBool(b)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		if field.OverflowInt(n) {
			return fmt.Errorf("overflow %s: %d", field.Type().Name(), n)
		}
		field.SetInt(n)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return err
		}
		if field.OverflowUint(n) {
			return fmt.Errorf("overflow %s: %d", field.Type().Name(), n)
		}
		field.SetUint(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(s, field.Type().Bits())
		if err != nil {
			return err
		}
		if field.OverflowFloat(n) {
			return fmt.Errorf("overflow %s: %g", field.Type().Name(), n)
		}
		field.SetFloat(n)
	}
	return nil
}

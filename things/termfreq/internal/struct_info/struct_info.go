package structinfo

import "reflect"

type StructInfo interface {
	Info(i interface{}) string
}

func Info(i interface{}) string {
	t := reflect.TypeOf(i)

	// If it's a pointer, get the element type
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}

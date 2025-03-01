package helpers

import (
	"clean-hex/pkg/framwork/helpers/is"
	"reflect"
)

func ToPtr(val reflect.Value) reflect.Value {
	typ := val.Type()
	if !is.Ptr(typ) {
		// this creates a pointer type inherently
		ptrVal := reflect.New(typ)
		ptrVal.Elem().Set(val)
		val = ptrVal
	}
	return val
}

func FromPtr(val reflect.Value) reflect.Value {
	if is.Ptr(val.Type()) {
		val = val.Elem()
	}
	return val
}

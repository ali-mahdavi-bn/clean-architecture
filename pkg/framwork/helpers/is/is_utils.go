package is

import (
	"clean-hex/pkg/framwork/adapter"
	"errors"
	"reflect"
)

func Empty(input any) bool {
	return input == nil
}

func Ptr(input any) bool {
	return reflect.ValueOf(input).Type().Kind() == reflect.Ptr
}

func EntityDeleted(input adapter.Entity) bool {
	return input != nil && input.IsDeleted()
}

func Error(err, target error) bool {
	return errors.Is(err, target)
}

func String(input any) bool {

	return reflect.ValueOf(input).Type().Kind() == reflect.String
}

func Bool(input any) bool {
	return reflect.ValueOf(input).Type().Kind() == reflect.Bool
}

func Int(input any) bool {
	switch reflect.ValueOf(input).Type().Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return true
	}
	return false
}

func Uint(input any) bool {
	switch reflect.ValueOf(input).Type().Kind() {
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return true
	}
	return false
}

func Float(input any) bool {
	switch reflect.ValueOf(input).Type().Kind() {
	case reflect.Float32,
		reflect.Float64:
		return true
	}
	return false
}

func Struct(input any) bool {
	return reflect.ValueOf(input).Type().Kind() == reflect.Struct
}

func Slice(input any) bool {
	return reflect.ValueOf(input).Type().Kind() == reflect.Slice
}

func Map(input any) bool {
	return reflect.ValueOf(input).Type().Kind() == reflect.Map
}

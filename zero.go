package pandati

import (
	"reflect"
)

var zeroStructMap = make(map[reflect.Type]interface{})

func IsZero[T any](v T) bool {
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return true
	}
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return true
	}
	if value.Kind() == reflect.Ptr && !value.IsNil() && value.Elem().Kind() == reflect.Struct {
		value = value.Elem()
	}
	zeroValue := reflect.Zero(value.Type())
	if value.Type().Comparable() {
		return reflect.DeepEqual(value.Interface(), zeroValue.Interface())
	}
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return value.Len() == 0
	case reflect.Struct:
		zeroStruct := zeroStructMap[value.Type()]
		if zeroStruct == nil {
			zeroStruct = reflect.New(value.Type()).Elem().Interface()
			zeroStructMap[value.Type()] = zeroStruct
		}
		return reflect.DeepEqual(value.Interface(), reflect.ValueOf(zeroStruct).Elem().Interface())
	}
	return false
}

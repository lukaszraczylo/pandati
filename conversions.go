package pandati

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

// StringToBytes converts a string to a byte slice without copying the underlying data.
func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// BytesToString converts a byte slice to a string without copying the underlying data.
func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// Convert any []byte reply to the desired reply type
// Useful for converting json into structs.
func ConvertReplyType[T any](desiredType T, reply []byte) T {
	if reply == nil {
		return interface{}(nil).(T)
	}
	desiredTypeType := reflect.TypeOf(desiredType)
	v := reflect.New(desiredTypeType).Interface()
	err := json.Unmarshal(reply, &v)
	if err != nil {
		return interface{}(nil).(T)
	}
	return reflect.ValueOf(v).Elem().Interface().(T)
}

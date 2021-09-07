package pandati

import (
	"fmt"
	"reflect"
)

// ExistsInSlice returns true if the given value exists in the given slice.
func ExistsInSlice(slice interface{}, value interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not a slice", slice))
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == value {
			return true
		}
	}
	return false
}

// RemoveFromSlice removes the given value from the given slice.
// The value is replaced with Zero value of the same type.
// For string it is an empty string.
// For integers it is 0.
func RemoveFromSlice(slice interface{}, value interface{}) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not a slice", slice))
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == value {
			s.Index(i).Set(reflect.Zero(s.Index(i).Type()))
			return
		}
	}
}

// RemoveFromSlice removes the given value from the given slice using provided index.
func RemoveFromSliceByIndex(slice interface{}, index int) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not a slice", slice))
	}
	if index < 0 || index >= s.Len() {
		panic(fmt.Sprintf("index %d is out of bounds", index))
	}
	s.Index(index).Set(reflect.Zero(s.Index(index).Type()))
}

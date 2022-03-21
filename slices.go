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
// The value is replaced with Zero value of the same type.
// For string it is an empty string.
// For integers it is 0.
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

// UniqueSlice returns slice with reordered and unique values.
// Please note that function does not modify the original slice but returns a new one.
func UniqueSlice(slice interface{}) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not a slice", slice))
	}

	for i := 0; i < s.Len(); i++ {
		for j := 0; j < s.Len(); j++ {
			if i == j {
				continue
			}
			if s.Index(i).Interface() == s.Index(j).Interface() {
				RemoveFromSliceByIndex(slice, j)
			}
		}
	}

	sliceCleaned := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
	for i := 0; i < s.Len(); i++ {
		if !IsZero(s.Index(i).Interface()) {
			sliceCleaned = reflect.Append(sliceCleaned, s.Index(i))
		}
	}
	fmt.Println(sliceCleaned.Interface())
	return sliceCleaned.Interface()
}

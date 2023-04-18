package pandati

import (
	"reflect"
)

var zeroStructMap = make(map[reflect.Type]interface{})

func IsZero[T comparable](v T) bool {
	var zero T
	return zero == v
}

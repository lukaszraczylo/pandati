package pandati

import (
	"math"
	"reflect"
)

func isZero(v reflect.Value) bool {
	if !v.IsValid() || v.IsZero() {
		return true
	}

	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return math.IsNaN(v.Float()) || math.Abs(v.Float()) < math.SmallestNonzeroFloat64

	case reflect.Complex64, reflect.Complex128:
		return math.IsNaN(real(v.Complex())) || math.IsNaN(imag(v.Complex())) || v.Complex() == 0

	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true

	case reflect.Slice, reflect.String, reflect.Map:
		return v.Len() == 0

	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true

	case reflect.Ptr:
		// If the value is a pointer, check if it is nil
		if v.IsNil() {
			return true
		}
		// If it is not nil, check if the pointed-to value is a zero value
		return isZero(v.Elem())

	default:
		return false
	}
}

// IsZero reports whether v is zero struct
// Does not support cycle pointers for performance, so as json
func IsZero(v interface{}) bool {
	return isZero(reflect.ValueOf(v))
}

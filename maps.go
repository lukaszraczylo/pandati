package pandati

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/imdario/mergo"
)

type FlattenOptions struct {
	Separator string
	Safe      bool
	MaxDepth  int
}

// Returns flat map from the provided nested map[string]interface{}
// For example:
//
// 		nested := map[string]interface{}{
// 			"a": map[string]interface{}{
// 				"b": map[string]interface{}{
// 					"c": "d",
// 				},
// 			},
// 		}
//
// 		flat := FlattenMap(nested)
//
// 		flat equals to map[string]interface{}{
// 			"a.b.c": "d",
// 		}
func FlattenMap(nested map[string]interface{}, opts *FlattenOptions) (m map[string]interface{}, err error) {
	if opts == nil {
		opts = &FlattenOptions{
			Separator: ".",
		}
	}
	m, err = flattenmap("", 0, nested, opts)
	return
}

func flattenmap(prefix string, depth int, nested interface{}, opts *FlattenOptions) (flatmap map[string]interface{}, err error) {
	flatmap = make(map[string]interface{})

	switch nested := nested.(type) {
	case map[string]interface{}:
		if opts.MaxDepth != 0 && depth >= opts.MaxDepth {
			flatmap[prefix] = nested
			return
		}
		if reflect.DeepEqual(nested, map[string]interface{}{}) {
			flatmap[prefix] = nested
			return
		}
		for k, v := range nested {
			// create new key
			newKey := k
			if prefix != "" {
				newKey = prefix + opts.Separator + newKey
			}
			fm1, fe := flattenmap(newKey, depth+1, v, opts)
			if fe != nil {
				err = fe
				return
			}
			updatemap(flatmap, fm1)
		}
	case []interface{}:
		if opts.Safe {
			flatmap[prefix] = nested
			return
		}
		if reflect.DeepEqual(nested, []interface{}{}) {
			flatmap[prefix] = nested
			return
		}
		for i, v := range nested {
			newKey := strconv.Itoa(i)
			if prefix != "" {
				newKey = prefix + opts.Separator + newKey
			}
			fm1, fe := flattenmap(newKey, depth+1, v, opts)
			if fe != nil {
				err = fe
				return
			}
			updatemap(flatmap, fm1)
		}
	default:
		flatmap[prefix] = nested
	}
	return
}

// update is the function that update to map with from
// example:
// to = {"hi": "there"}
// from = {"foo": "bar"}
// then, to = {"hi": "there", "foo": "bar"}
func updatemap(to map[string]interface{}, from map[string]interface{}) {
	for k, v := range from {
		to[k] = v
	}
}

// Unflatten the map, it returns a nested map of a map
// By default, the flatten has Separator equal to "."
//
// Example:
//
// 		flat := map[string]interface{}{
// 			"a.b.c": "d",
// 		}
//
// 		nested := UnflattenMap(flat)
//
// 		nested equals to map[string]interface{}{
// 			"a": map[string]interface{}{
// 				"b": map[string]interface{}{
// 					"c": "d",
// 				},
// 			},
// 		}
func UnflattenMap(flat map[string]interface{}, opts *FlattenOptions) (nested map[string]interface{}, err error) {
	if opts == nil {
		opts = &FlattenOptions{
			Separator: ".",
		}
	}
	nested, err = unflattenmap(flat, opts)
	return
}

func unflattenmap(flat map[string]interface{}, opts *FlattenOptions) (nested map[string]interface{}, err error) {
	nested = make(map[string]interface{})
	for k, v := range flat {
		n := updateunflattenmap(k, v, opts)
		err = mergo.Merge(&nested, n, mergo.WithOverride)
	}
	return
}

func updateunflattenmap(k string, v interface{}, opts *FlattenOptions) (n interface{}) {
	n = v
	keys := strings.Split(k, opts.Separator)
	for i := len(keys) - 1; i >= 0; i-- {
		temp := make(map[string]interface{})
		temp[keys[i]] = n
		n = temp
	}
	return
}

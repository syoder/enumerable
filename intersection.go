package enumerable

import (
	"reflect"
)

// Intersection returns a new slice containing only the items present
// in both slices
func Intersection(a, b AnySlice) AnySlice {
	mustBeSlice(a)
	mustBeSlice(b)

	aVal := reflect.ValueOf(a)
	bVal := reflect.ValueOf(b)
	aCount := aVal.Len()
	bCount := bVal.Len()
	output := makeFilterSlice(a, 0, aCount+bCount)
	keys := make(map[interface{}]bool)

	for i := 0; i < aCount; i++ {
		keys[aVal.Index(i).Interface()] = true
	}
	for i := 0; i < bCount; i++ {
		key := bVal.Index(i)
		if _, present := keys[key.Interface()]; present {
			output = reflect.Append(output, key)
		}
	}
	return output.Interface()
}

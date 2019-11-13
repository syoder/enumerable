package enumerable

import (
	"reflect"
)

// NotFound is returned by `Find` when no matching item is found
const NotFound = -1

// Find calls `f` for each item in the slice until `f` returns `true`. If an item
// is found, it returns the index of that item and the item itself. If none are
// found, it returns `NotFound, nil`.
func Find(slice AnySlice, f IterFunc) (int, AnyValue) {
	mustBeSlice(slice)
	mustBeIterFuncFilterForSlice(f, slice)

	input := reflect.ValueOf(slice)
	count := input.Len()
	funcValue := reflect.ValueOf(f)

	for i := 0; i < count; i++ {
		value := input.Index(i)
		funcOutputs := funcValue.Call([]reflect.Value{value})
		if funcOutputs[0].Bool() {
			return i, value.Interface()
		}
	}
	return NotFound, nil
}

// Any returns true if the given function returns true for any items in the slice
func Any(slice AnySlice, f IterFunc) bool {
	idx, _ := Find(slice, f)
	return idx != NotFound
}

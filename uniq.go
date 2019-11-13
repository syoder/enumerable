package enumerable

import (
	"reflect"
)

// Uniq returns a new slice with duplicates removed
func Uniq(slice AnySlice) AnySlice {
	mustBeSlice(slice)

	input := reflect.ValueOf(slice)
	count := input.Len()
	output := makeFilterSlice(slice, 0, count)
	keys := make(map[interface{}]bool)

	for i := 0; i < count; i++ {
		value := input.Index(i)
		if _, present := keys[value.Interface()]; !present {
			keys[value.Interface()] = true
			output = reflect.Append(output, value)
		}
	}
	return output.Interface()
}

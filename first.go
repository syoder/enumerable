package enumerable

import (
	"reflect"
)

// First copies the first n elements of a slice into a new slice
func First(slice AnySlice, n int) AnySlice {
	mustBeSlice(slice)

	input := reflect.ValueOf(slice)
	if input.Len() < n {
		n = input.Len()
	}
	output := makeFilterSlice(slice, n, n)

	for i := 0; i < n; i++ {
		output.Index(i).Set(input.Index(i))
	}
	return output.Interface()
}

// Last copies the last n elements of a slice into a new slice
func Last(slice AnySlice, n int) AnySlice {
	mustBeSlice(slice)

	input := reflect.ValueOf(slice)
	if input.Len() < n {
		n = input.Len()
	}
	output := makeFilterSlice(slice, n, n)

	for i := 0; i < n; i++ {
		j := input.Len() - n + i
		output.Index(i).Set(input.Index(j))
	}
	return output.Interface()
}

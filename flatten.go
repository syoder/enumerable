package enumerable

import (
	"reflect"
)

// Flatten returns a new slice by concatenating all the elements in `slice`
func Flatten(slice AnySlice) AnySlice {
	mustBeSlice(slice)
	input := reflect.ValueOf(slice)
	if input.Type().Elem().Kind() != reflect.Slice {
		panic("enumerable: Flatten requires a slice of slices")
	}
	elemType := input.Type().Elem().Elem()
	output := reflect.MakeSlice(reflect.SliceOf(elemType), 0, input.Len()*2)
	for i := 0; i < input.Len(); i++ {
		itemAry := input.Index(i)
		for j := 0; j < itemAry.Len(); j++ {
			output = reflect.Append(output, itemAry.Index(j))
		}
	}
	return output.Interface()
}

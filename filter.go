package enumerable

import (
	"reflect"
)

// Filter takes a slice of any type and a function which takes that type as
// a paramter and returns a bool to indicate whether that value should be
// included in the output.
func Filter(slice AnySlice, f IterFunc) AnySlice {
	mustBeSlice(slice)
	mustBeIterFuncFilterForSlice(f, slice)

	input := reflect.ValueOf(slice)
	funcValue := reflect.ValueOf(f)

	// find optimized filter func
	if impl := findFilterImpl(funcValue.Type()); impl != nil {
		// call non-generic / optimized func
		output := impl.Call(
			[]reflect.Value{input, funcValue},
		)[0]
		return output.Interface()
	}

	count := input.Len()
	output := makeFilterSlice(slice, 0, count)

	for i := 0; i < count; i++ {
		value := input.Index(i)
		funcOutputs := funcValue.Call([]reflect.Value{value})
		if funcOutputs[0].Bool() {
			output = reflect.Append(output, value)
		}
	}
	return output.Interface()
}

func findFilterImpl(typ reflect.Type) *reflect.Value {
	for _, f := range filterFuncs {
		if reflect.TypeOf(f).In(1) == typ {
			impl := reflect.ValueOf(f)
			return &impl
		}
	}
	return nil
}

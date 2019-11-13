package enumerable

import (
	"reflect"
)

// Map takes a slice of some type and a function which takes that type
// as a parameter, and returns a slice of the same type as the function returns
func Map(slice AnySlice, f IterFunc) AnySlice {
	mustBeSlice(slice)
	mustBeIterFuncForSlice(f, slice)

	input := reflect.ValueOf(slice)
	funcValue := reflect.ValueOf(f)

	// find optimized mapping func
	if impl := findMapImpl(funcValue.Type()); impl != nil {
		// call non-generic / optimized func
		output := impl.Call(
			[]reflect.Value{input, funcValue},
		)[0]
		return output.Interface()
	}

	// no optimized implementation found, use generic iteration via reflection
	count := input.Len()
	output := makeOutputSlice(f, count, count)
	args := make([]reflect.Value, 1)
	for i := 0; i < count; i++ {
		args[0] = input.Index(i)
		funcOutputs := funcValue.Call(args)
		output.Index(i).Set(funcOutputs[0])
	}
	return output.Interface()
}

func findMapImpl(typ reflect.Type) *reflect.Value {
	for _, mapFunc := range mapFuncs {
		if reflect.TypeOf(mapFunc).In(1) == typ {
			impl := reflect.ValueOf(mapFunc)
			return &impl
		}
	}
	return nil
}

func ToGenericSlice(in AnySlice) []AnyValue {
	return ConvertSlice(in, []AnyValue{}).([]AnyValue)
	// input := reflect.ValueOf(in)
	// count := input.Len()
	// out := make([]AnyValue, count, count)
	// for i := 0; i < count; i++ {
	// 	out[i] = input.Index(i).Interface()
	// }
	// return out
}

func ConvertSlice(in AnySlice, outType AnySlice) AnySlice {
	inVal := reflect.ValueOf(in)
	count := inVal.Len()
	output := reflect.MakeSlice(reflect.TypeOf(outType), count, count)
	for i := 0; i < count; i++ {
		output.Index(i).Set(inVal.Index(i))
	}
	return output.Interface()
}

// GenericMap does not use reflection, but does require the input slice to
// be of `[]AnyValue` type, and requires the function to do type assertions
// as necessary.
// var GenericMap = MapAnyValueToAnyValue
func GenericMap(in []AnyValue, f func(AnyValue) AnyValue) []AnyValue {
	return MapAnyValueToAnyValue(in, f)
}

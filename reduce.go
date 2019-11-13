package enumerable

import (
	"reflect"
)

// Reduce takes a slice of some type, a memo value of some other type, and a
// function which takes the memo type and the slice type as parameters and returns
// the memo type. Reduce returns the final memo.
func Reduce(slice AnySlice, memo AnyValue, f IterFunc) AnyValue {
	mustBeSlice(slice)
	mustBeIterFuncForSliceWithMemo(f, memo, slice)

	input := reflect.ValueOf(slice)
	funcValue := reflect.ValueOf(f)
	memoValue := reflect.ValueOf(memo)

	// find optimized mapping func
	if impl := findReduceImpl(funcValue.Type()); impl != nil {
		// call non-generic / optimized func
		output := impl.Call(
			[]reflect.Value{input, memoValue, funcValue},
		)[0]
		return output.Interface()
	}

	// no optimized implementation found, use generic iteration via reflection
	count := input.Len()
	for i := 0; i < count; i++ {
		funcOutputs := funcValue.Call(
			[]reflect.Value{
				memoValue,
				input.Index(i),
			},
		)
		memoValue = funcOutputs[0]
	}
	return memoValue.Interface()
}

func findReduceImpl(typ reflect.Type) *reflect.Value {
	for _, reduceFunc := range reduceFuncs {
		if reflect.TypeOf(reduceFunc).In(2) == typ {
			impl := reflect.ValueOf(reduceFunc)
			return &impl
		}
	}
	return nil
}

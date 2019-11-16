package enumerable

import (
	"fmt"
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

// SumAsInt sums all items in slice and returns value as int. Will panic
// if element values are floats
func SumAsInt(slice AnySlice) int {
	switch slice.(type) {
	case []int:
		return Reduce(slice, 0, func(memo, value int) int {
			return memo + value
		}).(int)
	case []int8, []int16, []int32, []int64, []uint8, []uint16, []uint32, []uint64, []uint:
		return SumAsInt(ConvertSlice(slice, []int{}))
	default:
		panic(fmt.Errorf("enumerable: cannot sum slice of type `%v` as int", reflect.TypeOf(slice)))
	}
}

// SumAsFloat64 will sum items in slice and return value as float64.
func SumAsFloat64(slice AnySlice) float64 {
	switch slice.(type) {
	case []int, []int8, []int16, []int32, []int64, []uint8, []uint16, []uint32, []uint64, []uint:
		return float64(SumAsInt(ConvertSlice(slice, []int{})))
	case []float64:
		return Reduce(slice, 0.0, func(memo, value float64) float64 {
			return memo + value
		}).(float64)
	case []float32:
		return SumAsFloat64(ConvertSlice(slice, []float64{}))
	default:
		panic(fmt.Errorf("enumerable: cannot sum slice of type `%v`", reflect.TypeOf(slice)))
	}
}

// Avg returns the average value of items in the slice
func Avg(slice AnySlice) float64 {
	count := reflect.ValueOf(slice).Len()
	return SumAsFloat64(slice) / float64(count)
}

package enumerable

import (
	"fmt"
	"reflect"
	"sort"
)

// SortBy sorts a slice by the return value of function which is given each item
// in the slice. The function's return value must be a `string` or `int`.
func SortBy(slice AnySlice, f IterFunc) {
	mustBeSlice(slice)
	fVal := reflect.ValueOf(f)
	if fVal.Kind() != reflect.Func {
		panic("enumerable: SortBy was given a non-function instead of a function")
	}

	sliceVal := reflect.ValueOf(slice)
	sort.Slice(slice, func(i, j int) bool {
		return less(
			fVal.Call([]reflect.Value{sliceVal.Index(i)})[0].Interface(),
			fVal.Call([]reflect.Value{sliceVal.Index(j)})[0].Interface(),
		)
	})
}

// Sort requires a slice with elements that are some sort of int or string
func Sort(slice AnySlice) {
	mustBeSlice(slice)

	sliceVal := reflect.ValueOf(slice)
	sort.Slice(slice, func(i, j int) bool {
		return less(
			sliceVal.Index(i).Interface(),
			sliceVal.Index(j).Interface(),
		)
	})
}

// Reverse reverses the order of items in the slice
func Reverse(slice AnySlice) {
	mustBeSlice(slice)
	count := reflect.ValueOf(slice).Len()
	swapper := reflect.Swapper(slice)
	i := 0
	j := count - 1
	for i < j {
		swapper(i, j)
		i++
		j--
	}
}

// Min returns the minimum value in `slice`
func Min(slice AnySlice) AnyValue {
	mustBeSlice(slice)
	count := reflect.ValueOf(slice).Len()
	if count == 0 {
		return nil
	}
	sliceVal := reflect.ValueOf(slice)
	var min AnyValue = sliceVal.Index(0).Interface()
	for i := 0; i < count; i++ {
		val := sliceVal.Index(i).Interface()
		if less(val, min) {
			min = val
		}
	}
	return min
}

// Max returns the maximum value in `slice`
func Max(slice AnySlice) AnyValue {
	mustBeSlice(slice)
	count := reflect.ValueOf(slice).Len()
	if count == 0 {
		return nil
	}
	sliceVal := reflect.ValueOf(slice)
	var max AnyValue = sliceVal.Index(0).Interface()
	for i := 0; i < count; i++ {
		val := sliceVal.Index(i).Interface()
		if less(max, val) {
			max = val
		}
	}
	return max
}

func less(a, b AnyValue) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case int8:
		return a.(int8) < b.(int8)
	case int16:
		return a.(int16) < b.(int16)
	case int32:
		return a.(int32) < b.(int32)
	case int64:
		return a.(int64) < b.(int64)
	case uint:
		return a.(uint) < b.(uint)
	case uint8:
		return a.(uint8) < b.(uint8)
	case uint16:
		return a.(uint16) < b.(uint16)
	case uint32:
		return a.(uint32) < b.(uint32)
	case uint64:
		return a.(uint64) < b.(uint64)
	case float64:
		return a.(float64) < b.(float64)
	case float32:
		return a.(float32) < b.(float32)
	case string:
		return a.(string) < b.(string)
	default:
		panic(fmt.Errorf("enumerable: cannot compare values of type %v", reflect.TypeOf(a)))
	}
}

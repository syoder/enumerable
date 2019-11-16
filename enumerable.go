package enumerable

import "reflect"

//go:generate sh -c "go run bin/generate_code.go < templates/map_funcs.go.tmpl > generated_map_funcs.go"
//go:generate sh -c "go run bin/generate_code.go < templates/reduce_funcs.go.tmpl > generated_reduce_funcs.go"
//go:generate sh -c "go run bin/generate_code.go < templates/filter_funcs.go.tmpl > generated_filter_funcs.go"

// Enumerable makes it easy to map, filter, reduce, etc using any slice
type Enumerable struct {
	Slice AnySlice
}

// New creates a new Enumerable struct
func New(slice AnySlice) Enumerable {
	return Enumerable{Slice: slice}
}

// GetSlice sets `slicePtr` to the stored slice
func (e Enumerable) GetSlice(slicePtr interface{}) {
	reflect.ValueOf(slicePtr).Elem().Set(reflect.ValueOf(e.Slice))
}

// ToGeneric converts the slice to be of type `[]AnyType`
func (e Enumerable) ToGeneric() Enumerable {
	return New(ToGenericSlice(e.Slice))
}

func (e Enumerable) Convert(sliceType AnySlice) Enumerable {
	return New(ConvertSlice(e.Slice, sliceType))
}

// Map iterates over the slice a returns a new Enumerable
func (e Enumerable) Map(f IterFunc) Enumerable {
	return New(Map(e.Slice, f))
}

// GenericMap does the same thing as `map` but requires a slice type
// of `[]AnyValue` and requires the function to do type assertions as
// necessary
func (e Enumerable) GenericMap(f func(AnyValue) AnyValue) Enumerable {
	return New(GenericMap(e.Slice.([]AnyValue), f))
}

// Filter returns a new Enumerable containing only the elements for which the given
// function returns `true`
func (e Enumerable) Filter(f IterFunc) Enumerable {
	return New(Filter(e.Slice, f))
}

// Reduce iterates over a slice passing the memo value into the function
func (e Enumerable) Reduce(memo AnyValue, f IterFunc) AnyValue {
	return Reduce(e.Slice, memo, f)
}

// ReduceToInt calls reduce but forces the memo to be an `int`
func (e Enumerable) ReduceToInt(memo int, f IterFunc) int {
	return e.Reduce(memo, f).(int)
}

// ReduceToFloat64 calls reduce but forces the memo to be a `float64`
func (e Enumerable) ReduceToFloat64(memo float64, f IterFunc) float64 {
	return e.Reduce(memo, f).(float64)
}

// SumAsInt returns the sum of items in the slice as an int. It will panic
// if the values are float values
func (e Enumerable) SumAsInt() int {
	return SumAsInt(e.Slice)
}

// SumAsFloat64 returns the sum of items in the slice as a float64 value
func (e Enumerable) SumAsFloat64() float64 {
	return SumAsFloat64(e.Slice)
}

// Avg returns the average value of items in the slice
func (e Enumerable) Avg() float64 {
	return Avg(e.Slice)
}

// First returns a new Enumerable containing only the first `n` elements
func (e Enumerable) First(n int) Enumerable {
	return New(First(e.Slice, n))
}

// Last returns a new Enumerable containing only the last `n` elements
func (e Enumerable) Last(n int) Enumerable {
	return New(Last(e.Slice, n))
}

// Find calls `f` for each item in the slice until `f` returns `true`. If an item
// is found, it returns the index of that item and the item itself. If none are
// found, it returns `NotFound, nil`.
func (e Enumerable) Find(f IterFunc) (int, AnyValue) {
	return Find(e.Slice, f)
}

// Any returns true if the given function returns true for any items in the slice
func (e Enumerable) Any(f IterFunc) bool {
	return Any(e.Slice, f)
}

// Uniq returns a new slice with duplicates removed
func (e Enumerable) Uniq() Enumerable {
	return New(Uniq(e.Slice))
}

// Flatten returns a new slice with each value of the slice concatenated
func (e Enumerable) Flatten() Enumerable {
	return New(Flatten(e.Slice))
}

// SortBy sorts an Enumerable by the return value of function which is given each item
// in the slice. The function's return value must be a `string` or `int`.
func (e Enumerable) SortBy(f IterFunc) Enumerable {
	SortBy(e.Slice, f)
	return e
}

func (e Enumerable) Sort() Enumerable {
	Sort(e.Slice)
	return e
}

// Reverse reverses the order of the items in the slice
func (e Enumerable) Reverse() Enumerable {
	Reverse(e.Slice)
	return e
}

// Min returns the minimum value in the slice
func (e Enumerable) Min() AnyValue {
	return Min(e.Slice)
}

// Max returns the maximum value in the slice
func (e Enumerable) Max() AnyValue {
	return Max(e.Slice)
}

func (e Enumerable) Group(keyFunc IterFunc) AnyValue {
	return Group(e.Slice, keyFunc)
}

// JoinAsString returns a string with each item converted to string
// and joined with `sep`.
func (e Enumerable) JoinAsString(sep string) string {
	return JoinAsString(e.Slice, sep)
}

// JoinAsStringWithFormat returns a string with each item converted to string
// using the given format string (see fmt.Printf) and joined with `sep`.
func (e Enumerable) JoinAsStringWithFormat(format, sep string) string {
	return JoinAsStringWithFormat(e.Slice, format, sep)
}

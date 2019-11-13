package enumerable

// FilterInts returns a slice of int where `f` returns true
func FilterInts(in []int, f func(int) bool) []int {
	out := make([]int, 0, len(in))
	for _, value := range in {
		if f(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterFloat64s returns a slice of float64 where `f` returns true
func FilterFloat64s(in []float64, f func(float64) bool) []float64 {
	out := make([]float64, 0, len(in))
	for _, value := range in {
		if f(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterBools returns a slice of bool where `f` returns true
func FilterBools(in []bool, f func(bool) bool) []bool {
	out := make([]bool, 0, len(in))
	for _, value := range in {
		if f(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterStrings returns a slice of string where `f` returns true
func FilterStrings(in []string, f func(string) bool) []string {
	out := make([]string, 0, len(in))
	for _, value := range in {
		if f(value) {
			out = append(out, value)
		}
	}
	return out
}

// FilterAnyValues returns a slice of AnyValue where `f` returns true
func FilterAnyValues(in []AnyValue, f func(AnyValue) bool) []AnyValue {
	out := make([]AnyValue, 0, len(in))
	for _, value := range in {
		if f(value) {
			out = append(out, value)
		}
	}
	return out
}

var filterFuncs = []interface{}{
	FilterInts,
	FilterFloat64s,
	FilterBools,
	FilterStrings,
	FilterAnyValues,
}
